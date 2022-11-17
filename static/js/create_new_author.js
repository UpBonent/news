/* function handleSubmit(event) {
    event.preventDefault();

    const data = new FormData(event.target);

    const value = data.get('name');

    console.log(value);
}

const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit);

 */


const options = {
    hostname: 'localhost:8080',
    path: '',
    method: 'GET', // default
}
const req = http.request(options, res => {
    console.log(res.statusCode);
});
req.end();