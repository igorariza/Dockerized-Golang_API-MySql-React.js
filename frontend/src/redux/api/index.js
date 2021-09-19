import axios from "axios";

const baseUrl = 'https://www.easports.com/fifa/ultimate-team/api/fut/item';

export const apiCall = (url, data, headers, method) => axios ({
    method,
    url: baseUrl +url,
    headers
});