// 使用md5编写encryptPassword函数，用于加密密码
import md5 from 'js-md5'

export function encryptPassword(password) {
    return md5(password)
}
