import http from 'k6/http';
import { check, sleep } from 'k6';
import { Trend } from 'k6/metrics';

// Base URL
const BASE_URL = `http://${__ENV.APP_NAME || 'bbb-voting-app'}:${__ENV.APP_PORT || 8081}`;

// Metrics for CAPTCHA
const generateCaptchaTrend = new Trend('generate_captcha_duration');
const serveCaptchaTrend = new Trend('serve_captcha_duration');
const validateCaptchaTrend = new Trend('validate_captcha_duration');

// Stages and thresholds configuration
export let options = {
    stages: [
        { duration: '1m', target: 100 }, // Ramp-up to 100 VUs in 1 minute
        { duration: '3m', target: 100 }, // Stay at 100 VUs for 3 minutes
        { duration: '1m', target: 0 },   // Ramp-down to 0 VUs in 1 minute
    ],
    thresholds: {
        // Thresholds for CAPTCHA
        'generate_captcha_duration': ['p(95)<500'], // 95% of requests must complete below 500ms
        'serve_captcha_duration': ['p(95)<500'],
        'validate_captcha_duration': ['p(95)<500'],
    },
};

// CAPTCHA test
export default function testCaptcha() {
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