import http from 'k6/http';
import { check, sleep } from 'k6';

const APP_NAME = __ENV.APP_NAME || 'bbb-voting-app';
const APP_PORT = __ENV.APP_PORT || 8081;
const BASE_URL = `http://${APP_NAME}:${APP_PORT}`;

export const options = {
    stages: [
        { duration: '1m', target: 100 }, // ramp up to 100 users
        { duration: '3m', target: 100 }, // stay at 100 users for 3 minutes
        { duration: '1m', target: 0 },   // ramp down to 0 users
    ],
    teardownTimeout: '30s', // Allow time for teardown
};

// Test Generate CAPTCHA endpoint
export function testGenerateCaptcha() {
    const generatedCaptcha = http.get(`${BASE_URL}/v1/generate-captcha`);
    check(generatedCaptcha, { 'status was 200': (r) => r.status == 200 });

    if (generatedCaptcha.status !== 200) {
        console.error(`Failed to generate CAPTCHA: ${generatedCaptcha.status}`);
        return;
    }

    sleep(1);

    // Test Serve CAPTCHA endpoint
    const captchaID = generatedCaptcha.json().id;

    const serveCaptcha = http.get(`${BASE_URL}/v1/captcha/${captchaID}`);
    check(serveCaptcha, { 'status was 200': (r) => r.status == 200 });

    if (serveCaptcha.status !== 200) {
        console.error(`Failed to serve CAPTCHA: ${serveCaptcha.status}`);
        return;
    }

    sleep(1);

    // Test Validate CAPTCHA endpoint
    const validateCaptchaPayload = JSON.stringify({
        captcha_id: captchaID,
        captcha_solution: 'test_solution' // Use the test solution
    });
    const validatedCaptcha = http.post(`${BASE_URL}/v1/validate-captcha`, validateCaptchaPayload, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(validatedCaptcha, { 'status was 200': (r) => r.status == 200 });

    if (validatedCaptcha.status !== 200) {
        console.error(`Failed to validate CAPTCHA: ${validatedCaptcha.status}`);
        return;
    }

    sleep(1);
}

// Test Create Participant endpoint
export function testCreateParticipant() {
    const participantPayload = JSON.stringify({
        name: 'Participant 1',
        age: 25,
        gender: 'gender'
    });
    const firstCreatedParticipant = http.post(`${BASE_URL}/v1/participants`, participantPayload, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(firstCreatedParticipant, { 'status was 201': (r) => r.status == 201 });
    sleep(1);

    // Test Get Participants endpoint
    const participants = http.get(`${BASE_URL}/v1/participants`);
    check(participants, { 'status was 200': (r) => r.status == 200 });
    sleep(1);

    // Test Get Participant endpoint
    const participantID = firstCreatedParticipant.json().id;
    const participant = http.get(`${BASE_URL}/v1/participants/${participantID}`);
    check(participant, { 'status was 200': (r) => r.status == 200 });
    sleep(1);

    // Test Delete Participant endpoint
    const response = http.del(`${BASE_URL}/v1/participants/${participantID}`);
    check(response, { 'status was 204': (r) => r.status == 204 });
    sleep(1);
}

// Test Vote endpoint
export function testVote() {
    // Step 1: Create participants for voting
    const participantPayload1 = JSON.stringify({
        name: 'Participant 4',
        age: 25,
        gender: 'gender'
    });
    const participantPayload2 = JSON.stringify({
        name: 'Participant 5',
        age: 30,
        gender: 'gender'
    });

    const createdParticipant1 = http.post(`${BASE_URL}/v1/participants`, participantPayload1, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(createdParticipant1, { 'status was 201': (r) => r.status == 201 });

    const createdParticipant2 = http.post(`${BASE_URL}/v1/participants`, participantPayload2, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(createdParticipant2, { 'status was 201': (r) => r.status == 201 });

    // Refresh the list of participants
    sleep(1); // Wait for the data to propagate
    const participantsResponse = http.get(`${BASE_URL}/v1/participants`);
    check(participantsResponse, { 'status was 200': (r) => r.status == 200 });
    const participants = participantsResponse.json();

    // Step 2: Generate CAPTCHA
    const generatedCaptcha = http.get(`${BASE_URL}/v1/generate-captcha`);
    check(generatedCaptcha, { 'status was 200': (r) => r.status == 200 });

    if (generatedCaptcha.status !== 200) {
        console.error(`Failed to generate CAPTCHA: ${generatedCaptcha.status}`);
        return;
    }

    sleep(1);

    // Step 3: Validate CAPTCHA
    const captchaID = generatedCaptcha.json().captcha_id || generatedCaptcha.json().id;
    const validateCaptchaPayload = JSON.stringify({
        captcha_id: captchaID,
        captcha_solution: 'test_solution' // Use the test solution
    });
    const validatedCaptcha = http.post(`${BASE_URL}/v1/validate-captcha`, validateCaptchaPayload, {
        headers: { 'Content-Type': 'application/json' }
    });
    check(validatedCaptcha, { 'status was 200': (r) => r.status == 200 });

    if (validatedCaptcha.status !== 200) {
        console.error(`Failed to validate CAPTCHA: ${validatedCaptcha.status}`);
        return;
    }

    sleep(1);

    // Extract CAPTCHA token from response headers
    const captchaToken = validatedCaptcha.headers['X-Captcha-Token'];

    // Step 4: Cast Vote
    const randomIndex = Math.floor(Math.random() * participants.length);
    const participantID = participants[randomIndex].id; // Vote for a random participant
    const castVotePayload = JSON.stringify({
        participant_id: participantID,
        voter_id: 'some-voter-id',
        ip_address: '127.0.0.1',
        user_agent: 'k6-load-test',
        region: 'some-region',
        device: 'some-device'
    });
    const castVoteRes = http.post(`${BASE_URL}/v1/votes`, castVotePayload, {
        headers: {
            'Content-Type': 'application/json',
            'X-Captcha-Token': captchaToken
        }
    });
    check(castVoteRes, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}

// Test Get Results endpoint
export function testGetResults() {
    // Test Get Partial Results endpoint
    const partialResults = http.get(`${BASE_URL}/v1/results/partial`);
    check(partialResults, { 'status was 200': (r) => r.status == 200 });
    sleep(1);

    // Test Get Final Results endpoint
    const finalResults = http.get(`${BASE_URL}/v1/results/final`);
    check(finalResults, { 'status was 200': (r) => r.status == 200 });
    sleep(1);
}

// Main function to run all tests
export default function () {
    testGenerateCaptcha();
    testCreateParticipant();
    testVote();
    testGetResults();
}