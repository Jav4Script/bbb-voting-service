import http from 'k6/http';
import { check } from 'k6';
import { Trend } from 'k6/metrics';

// Base URL
const BASE_URL = `http://${__ENV.APP_NAME || 'bbb-voting-app'}:${__ENV.APP_PORT || 8081}`;

// Metrics
const finalResultsTrend = new Trend('final_results_duration');

// Stages and thresholds configuration
export let options = {
    stages: [
        { duration: '30s', target: 35 },  // Ramp-up to 35 VUs in 30 seconds
        { duration: '2m', target: 35 },   // Stay at 35 VUs for 2 minutes
        { duration: '30s', target: 0 },   // Ramp-down to 0 VUs in 30 seconds
    ],
    thresholds: {
        'final_results_duration': ['p(95)<750'], // 95% of requests must complete below 750ms
    },
};

// Main test function
export default function () {
    // Get final results
    const finalResults = http.get(`${BASE_URL}/v1/results/final`);
    check(finalResults, { 'Final results retrieved successfully': (r) => r.status === 200 });
    finalResultsTrend.add(finalResults.timings.duration);
}