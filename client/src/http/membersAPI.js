import axios from "axios"

const $host = axios.create({
    baseURL: "http://localhost:8080"
})
export const create =  async( email, name) => {
    return await $host.post('/member', { email,name})
}
export const get =  async() => {
    return  await $host.get('/members', )
}