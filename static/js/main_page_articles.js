fetch('/articles')
    .then(data => data.json())
    .then(data => {
        showArticles(data)
    })

let articleSection = document.querySelector('section')

function showArticles(jsonObj) {
    let article = document.createElement('article');
    let list = document.createElement('ul');

    for (let i = 0; i < jsonObj.length; i++) {
        let listItem = document.createElement('li')
        let h3 = document.createElement('h3');
        let p1 = document.createElement('p');
        let time = document.createElement('time');

        h3.textContent = jsonObj[i].header;
        p1.textContent = jsonObj[i].text;
        time.textContent = 'Date of create: ' + jsonObj[i].date_create;

        listItem.appendChild(h3);
        listItem.appendChild(p1);
        listItem.appendChild(time);

        list.appendChild(listItem)
    }
    article.appendChild(list)
    articleSection.appendChild(article)
}