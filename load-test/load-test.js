import http from 'k6/http';
import { check, sleep } from 'k6';
import { Trend } from 'k6/metrics';

// Base URL
const BASE_URL = `http://${__ENV.APP_NAME || 'bbb-voting-app'}:${__ENV.APP_PORT || 8081}`;

// Metrics
const generateCaptchaTrend = new Trend('generate_captcha_duration');
const validateCaptchaTrend = new Trend('validate_captcha_duration');
const serveCaptchaTrend = new Trend('serve_captcha_duration');
const castVoteTrend = new Trend('cast_vote_duration');
const partialResultsTrend = new Trend('partial_results_duration');
const createParticipantTrend = new Trend('create_participant_duration');
const getParticipantsTrend = new Trend('get_participants_duration');
const getParticipantTrend = new Trend('get_participant_duration');
const deleteParticipantTrend = new Trend('delete_participant_duration');
const finalResultsTrend = new Trend('final_results_duration');

// Test options
export let options = {
    stages: [
        { duration: '30s', target: 35 },
        { duration: '2m', target: 35 },
        { duration: '30s', target: 0 },
    ],
    thresholds: {
        'generate_captcha_duration': ['p(95)<500'],
        'validate_captcha_duration': ['p(95)<500'],
        'serve_captcha_duration': ['p(95)<500'],
        'cast_vote_duration': ['p(95)<500'],
        'partial_results_duration': ['p(95)<500'],
        'create_participant_duration': ['p(95)<500'],
        'get_participants_duration': ['p(95)<500'],
        'get_participant_duration': ['p(95)<500'],
        'delete_participant_duration': ['p(95)<500'],
        'final_results_duration': ['p(95)<750'],
        http_req_duration: ['p(95)<500'],
        http_req_failed: ['rate<0.01'],
    },
};

// Setup function to create participants and return their IDs
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
    testParticipants();
    testCaptcha();
    testVote(data);
    testGetResults();
}

function testParticipants() {
    const participantPayload = JSON.stringify({
        name: `Participant ${__VU}-${__ITER}`,
        age: Math.floor(Math.random() * 50) + 18,
        gender: 'gender'
    });

    let response;

    // Create participant
    response = http.post(`${BASE_URL}/v1/participants`, participantPayload, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(response, { 'status was 201': (r) => r.status == 201 });
    createParticipantTrend.add(response.timings.duration);
    sleep(1);

    // Get all participants
    response = http.get(`${BASE_URL}/v1/participants`);
    check(response, { 'status was 200': (r) => r.status == 200 });
    getParticipantsTrend.add(response.timings.duration);
    sleep(1);

    // Get specific participant
    const participantID = response.json()[0].id;
    response = http.get(`${BASE_URL}/v1/participants/${participantID}`);
    check(response, { 'status was 200': (r) => r.status == 200 });
    getParticipantTrend.add(response.timings.duration);
    sleep(1);

    // Delete participant
    response = http.del(`${BASE_URL}/v1/participants/${participantID}`);
    check(response, { 'status was 204': (r) => r.status == 204 });
    deleteParticipantTrend.add(response.timings.duration);
    sleep(1);
}

function testCaptcha() {
    // Generate CAPTCHA
    const generatedCaptcha = http.get(`${BASE_URL}/v1/generate-captcha`);
    check(generatedCaptcha, { 'status was 200': (r) => r.status == 200 });
    generateCaptchaTrend.add(generatedCaptcha.timings.duration);

    if (generatedCaptcha.status !== 200) {
        console.error(`Failed to generate CAPTCHA: ${generatedCaptcha.status}`);
        return;
    }

    sleep(1);

    // Serve CAPTCHA
    const captchaID = generatedCaptcha.json().id;
    const serveCaptcha = http.get(`${BASE_URL}/v1/captcha/${captchaID}`);
    check(serveCaptcha, { 'status was 200': (r) => r.status == 200 });
    serveCaptchaTrend.add(serveCaptcha.timings.duration);

    if (serveCaptcha.status !== 200) {
        console.error(`Failed to serve CAPTCHA: ${serveCaptcha.status}`);
        return;
    }

    sleep(1);

    // Validate CAPTCHA
    const validateCaptchaPayload = JSON.stringify({
        captcha_id: captchaID,
        captcha_solution: 'test_solution'
    });
    const validatedCaptcha = http.post(`${BASE_URL}/v1/validate-captcha`, validateCaptchaPayload, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(validatedCaptcha, { 'status was 200': (r) => r.status == 200 });
    validateCaptchaTrend.add(validatedCaptcha.timings.duration);

    if (validatedCaptcha.status !== 200) {
        console.error(`Failed to validate CAPTCHA: ${validatedCaptcha.status}`);
        return;
    }

    sleep(1);
}

function testVote(data) {
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

function testGetResults() {
    // Get final results
    const finalResults = http.get(`${BASE_URL}/v1/results/final`);
    check(finalResults, { 'Final results retrieved successfully': (r) => r.status === 200 });
    finalResultsTrend.add(finalResults.timings.duration);
}