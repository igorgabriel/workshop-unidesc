import axios from 'axios';
import url from 'url';

const workshopAPI = axios.create({
    baseURL: url.resolve('http://localhost:8888', '/v1'),
    withCredentials: true
});

const getWorkshops = (onSuc, onErr) => {
    workshopAPI.get('/workshops').then((response) => onSuc(response.data.workshops))
        .catch((error) => onErr(getError(error)));
}

const getWorkshopById = (id, onSuc, onErr) => {
    workshopAPI.get(`/workshops/${id}`).then((response) => onSuc(response.data.workshop))
        .catch((error) => onErr(getError(error)));
}

const saveWorkshop = (workshop, onSuc, onErr) => {
    workshopAPI.post('/workshops', workshop).then((response) => onSuc(response.data))
        .catch((error) => onErr(getError(error)));
}

const updateWorkshop = (workshop, onSuc, onErr) => {
    const id = workshop.id;
    delete workshop.id;
    workshopAPI.put(`/workshops/${id}`, workshop).then((response) => onSuc(response.data))
        .catch((error) => onErr(getError(error)));
}

const deleteWorkshop = (id, onSuc, onErr) => {
    workshopAPI.delete(`/workshops/${id}`)
        .then((response) => onSuc(response.data))
        .catch((error) => onErr(getError(error)));
}

const getError = (error) => {
    if (error) {
        if (error.response) {
            if (error.response.data) {
                return error.response.data
            }
            return error.response
        }
        return error
    }
}

export {
    getWorkshops, getWorkshopById, saveWorkshop, updateWorkshop, deleteWorkshop
};
export default workshopAPI;
