import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '10s', target: 100 }, // below normal load
    { duration: '1m', target: 100 },
    { duration: '10s', target: 1400 }, // spike to 1400 users
    { duration: '3m', target: 1400 }, // stay at 1400 for 3 minutes
    { duration: '10s', target: 100 }, // scale down. Recovery stage.
    { duration: '3m', target: 100 },
    { duration: '10s', target: 0 },
  ],
};
export default function () {
    const hosts = __ENV.HOSTNAMES.split(';');
    const tags = __ENV.TAGS.split(';');
    for (let i = 0; i < hosts.length-1; i++) {
      const res = http.get(`http://${hosts[i]}`, { tags: { my_custom_tag: tags[i] } });
      sleep(1);
    }
}