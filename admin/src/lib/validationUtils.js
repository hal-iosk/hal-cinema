import validator from 'validator';

const ValidationUtils = {

  isEmail(email) {
    return validator.isEmail(email);
  },

  isBlank(text) {
    return text === "" || text === null
  },

  isNumber(number) {
    return validator.isInt(number)
  },

  isURL(url) {
    return validator.isURL(url)
  },

  pushMessage(message, vue) {
    vue.$toast.open({
      message,
      type: "is-danger"
    })
  }

}

export default ValidationUtils;