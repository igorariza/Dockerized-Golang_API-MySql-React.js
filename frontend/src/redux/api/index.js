import axios from "axios";

const baseUrl = 'http://localhost:8080/';

export const apiCall = (url, data, headers, method) => axios ({
    method,
    url: baseUrl + url,
    headers
});