import http from 'k6/http';
import { check } from 'k6';
import { Trend } from 'k6/metrics';

// Base URL
const BASE_URL = `http://${__ENV.APP_NAME || 'bbb-voting-app'}:${__ENV.APP_PORT || 8081}`;

// Metrics
const generateCaptchaTrend = new Trend('generate_captcha_duration');
const validateCaptchaTrend = new Trend('validate_captcha_duration');
const castVoteTrend = new Trend('cast_vote_duration');
const partialResultsTrend = new Trend('partial_results_duration');

// Test options
export let options = {
    scenarios: {
        constant_rps: {
            executor: 'constant-arrival-rate',
            rate: 1000, // 1000 requests per second
            timeUnit: '1s', // RPS is defined per second
            duration: '1m', // Test duration is 1 minute
            preAllocatedVUs: 200, // Pre-allocate 200 VUs
            maxVUs: 500, // Allow up to 500 VUs
        },
    },
    thresholds: {
        // Thresholds for CAPTCHA-related operations
        'generate_captcha_duration': ['p(95)<500'],
        'validate_captcha_duration': ['p(95)<500'],

        // Threshold for voting operation
        'cast_vote_duration': ['p(95)<500'],

        // Threshold for partial results operation
        'partial_results_duration': ['p(95)<500'],

        // Global thresholds
        http_req_duration: ['p(95)<500'], // 95% of all requests must complete below 500ms
        http_req_failed: ['rate<0.01'],  // Less than 1% of requests should fail
    },
};

// Setup: Create participants once before the test starts
export function setup() {
    // Create participant 1
    const participantPayload1 = JSON.stringify({
        name: 'Participant 1',
        age: 25,
        gender: 'gender',
    });
    const createdParticipant1 = http.post(`${BASE_URL}/v1/participants`, participantPayload1, {
        headers: { 'Content-Type': 'application/json' },
    });
    check(createdParticipant1, { 'Participant 1 created successfully': (r) => r.status === 201 });

    // Create participant 2
    const participantPayload2 = JSON.stringify({
        name: 'Participant 2',
        age: 30,
        gender: 'gender',
    });
    const createdParticipant2 = http.post(`${BASE_URL}/v1/participants`, participantPayload2, {
        headers: { 'Content-Type': 'application/json' },
    });
    check(createdParticipant2, { 'Participant 2 created successfully': (r) => r.status === 201 });

    // Return participant IDs for use in the test
    return {
        participant1Id: createdParticipant1.json().id,
        participant2Id: createdParticipant2.json().id,
    };
}

// Main test function
export default function (data) {
    // Select a random participant for voting
    const participantId = Math.random() < 0.5 ? data.participant1Id : data.participant2Id;

    // Generate CAPTCHA
    const generatedCaptcha = http.get(`${BASE_URL}/v1/generate-captcha`);
    check(generatedCaptcha, { 'CAPTCHA generated successfully': (r) => r.status === 200 });
    generateCaptchaTrend.add(generatedCaptcha.timings.duration);

    if (generatedCaptcha.status !== 200) {
        console.error(`Failed to generate CAPTCHA: ${generatedCaptcha.status}`);
        return;
    }

    const captchaID = generatedCaptcha.json().captcha_id || generatedCaptcha.json().id;

    // Validate CAPTCHA
    const validateCaptchaPayload = JSON.stringify({
        captcha_id: captchaID,
        captcha_solution: 'test_solution',
    });
    const validatedCaptcha = http.post(`${BASE_URL}/v1/validate-captcha`, validateCaptchaPayload, {
        headers: { 'Content-Type': 'application/json' },
    });
    check(validatedCaptcha, { 'CAPTCHA validated successfully': (r) => r.status === 200 });
    validateCaptchaTrend.add(validatedCaptcha.timings.duration);

    if (validatedCaptcha.status !== 200) {
        console.error(`Failed to validate CAPTCHA: ${validatedCaptcha.status}`);
        return;
    }

    const captchaToken = validatedCaptcha.headers['X-Captcha-Token'];

    // Cast vote
    const castVotePayload = JSON.stringify({
        participant_id: participantId,
        voter_id: 'some-voter-id',
        ip_address: '127.0.0.1',
        user_agent: 'k6-load-test',
        region: 'some-region',
        device: 'some-device',
    });
    const castVoteRes = http.post(`${BASE_URL}/v1/votes`, castVotePayload, {
        headers: {
            'Content-Type': 'application/json',
            'X-Captcha-Token': captchaToken,
        },
    });
    check(castVoteRes, { 'Vote cast successfully': (r) => r.status === 200 });
    castVoteTrend.add(castVoteRes.timings.duration);

    // Get partial results
    const partialResults = http.get(`${BASE_URL}/v1/results/partial`);
    check(partialResults, { 'Partial results retrieved successfully': (r) => r.status === 200 });
    partialResultsTrend.add(partialResults.timings.duration);
}