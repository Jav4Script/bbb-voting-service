import testCaptcha from './load-test-captcha.js';
import testParticipants from './load-test-participant.js';
import testVote from './load-test-vote.js';
import testGetResults from './load-test-results.js';

export default function () {
    testParticipants();
    testCaptcha();
    testVote();
    testGetResults();
}