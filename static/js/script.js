const body = document.querySelector('body'),
      sidebar = body.querySelector('.side-bar'),
      toggle = body.querySelector('.toggle'),
      switcher = body.querySelector('.switcher'),
      moodeText = body.querySelector('.text-day'),
      main = body.querySelector('.main'),
      image = body.querySelector('.img'),
      button = body.querySelector('.btn'),
      url = 'https://dog.ceo/api/breeds/image/random';

body.addEventListener('click', menu)

function menu (event) {
   if (event.target.closest('.toggle')) {
        sidebar.classList.toggle('close');
   }
   if (event.target.closest('.main')) {
        sidebar.classList.add('close');
   }
};

switcher.addEventListener('click', () =>{

    if (body.classList.contains('light')){
        body.classList.remove('light');
        body.classList.toggle('dark');
        moodeText.innerText = 'Light way';
    }else {
        body.classList.remove('dark');
        body.classList.toggle('light');
        moodeText.innerText = 'Dark side';
    }
   
});

async function fetchHandler() {
    try {
        const response = await fetch(url);
        const data = await response.json();
        image.src = data.message;
    } catch (error) {
        console.log(error);
    }
}

button.addEventListener('click', () => {
    let isloaded = image.complete;

    if (isloaded) {
        fetchHandler();
    };
})