import axios from "axios";
import { change_auth } from "../../redux/features/userSlice";
import { useDispatch } from "react-redux";

export const API_URL = "http://10.0.0.59:8080/api/v1";

// function checkAuth(){
//     axios({
//        method:'get',
//        url: `${API_URL}/resfresh`,
//        withCredentials: true
//     })
//     .then((result)=>{
//         const {data} = result;
//         localStorage.setItem('token', data.access_token)
//         change_auth(true)
//     })
   
// }



const $api = axios.create({
    withCredentials: true,
    baseURL: API_URL,
})

$api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`
    return config;
})

$api.interceptors.response.use((config) =>{
    return config;
}, (error) =>{
    const originalRequest = error.config
    if (error.response.status === 401 && error.config && !originalRequest._isRetry){
        originalRequest._isRetry = true
        try{
            axios({
                method:'get',
                url: `${API_URL}/user/refresh`,
                withCredentials: true
             })
             .then((result)=>{
                const {data} = result;
                localStorage.setItem('token', data.access_token)
             })
             return $api.request(originalRequest)
        } catch(e) {
            console.log("ВЫ НЕ АВТОРИЗОВАНЫ")
        }
        
    }
    throw error;
})

export default $api;