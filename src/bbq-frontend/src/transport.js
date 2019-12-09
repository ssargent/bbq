import axios from "axios";

const transport = axios.create({
    withCredentials: true
  });
  
  (function() {
    const token = localStorage.getItem("bbq-authenticated");

    if (token) {
        const tokenObj = JSON.parse(token);
        console.log(tokenObj);
        axios.defaults.headers.common['Authorization'] = `Bearer ${tokenObj.token}`;
    } else {
        axios.defaults.headers.common['Authorization'] = null;
        /*if setting null does not remove `Authorization` header then try     
          delete axios.defaults.headers.common['Authorization'];
        */
    }
})();

export { transport }