import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  const str = __ENV.HOSTNAMES.split(',');
  for (let i = 0; i < str.length-1; i++) {
    const res = http.get(`http://${str[i]}`);
    sleep(1);
  }
}