import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  const hosts = __ENV.HOSTNAMES.split(';');
  const tags = __ENV.TAGS.split(';');
  for (let i = 0; i < hosts.length-1; i++) {
    const res = http.get(`http://${hosts[i]}`, { tags: { my_custom_tag: tags[i] } });
    sleep(1);
  }
}