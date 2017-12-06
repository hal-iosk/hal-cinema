import axios from 'axios'

class HttpUtils {

  Login(email, password) {
    return axios.post("/api", {
      email,
      password
    })
  }

}

export default new HttpUtils;
