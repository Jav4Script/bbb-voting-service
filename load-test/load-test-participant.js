import http from 'k6/http';
import { check, sleep } from 'k6';
import { Trend } from 'k6/metrics';

const BASE_URL = `http://${__ENV.APP_NAME || 'bbb-voting-app'}:${__ENV.APP_PORT || 8081}`;
const createParticipantTrend = new Trend('create_participant_duration');
const getParticipantsTrend = new Trend('get_participants_duration');
const getParticipantTrend = new Trend('get_participant_duration');
const deleteParticipantTrend = new Trend('delete_participant_duration');



export let options = {
    stages: [
        { duration: '1m', target: 100 }, // Ramp-up to 100 VUs over 1 minute
        { duration: '3m', target: 100 }, // Stay at 100 VUs for 3 minutes
        { duration: '1m', target: 0 },  // Ramp-down to 0 VUs over 1 minute
    ],
    thresholds: {
        'create_participant_duration': ['p(95)<500'], // 95% of requests must complete below 500ms
        'get_participants_duration': ['p(95)<500'],
        'get_participant_duration': ['p(95)<500'],
        'delete_participant_duration': ['p(95)<500'],
    },
};

export default function testParticipants() {
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