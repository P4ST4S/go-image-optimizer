import http from 'k6/http';
import { check } from 'k6';

// Configuration du fichier à uploader
// "b" pour binaire
const binFile = open('./test.jpg', 'b'); 

export const options = {
  // On monte progressivement la charge
  stages: [
    { duration: '5s', target: 20 },  // On monte à 20 users en 5s
    { duration: '10s', target: 50 }, // On tient 50 users pendant 10s
    { duration: '5s', target: 0 },   // On redescend
  ],
};

export default function () {
  const url = 'http://localhost:8080/upload';
  
  const data = { 
    image: http.file(binFile, 'test.jpg'), // Simule le form-data
  };

  // Envoie la requête
  const res = http.post(url, data);

  // Vérifie que le serveur a répondu 200 OK
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
}