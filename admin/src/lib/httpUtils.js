import axios from 'axios'

class HttpUtils {

  Login(email, password) {
    return axios.post("/api/admin/signin", {
      email,
      password
    })
  }

}

export default new HttpUtils;
