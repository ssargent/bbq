import axios from "axios";

const transport = axios.create({
  withCredentials: true
});

export { transport };
