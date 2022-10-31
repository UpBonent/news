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
        let h4 = document.createElement('h4');
        let p1 = document.createElement('p');
        let p2 = document.createElement('p');

        h4.textContent = jsonObj[i].header;
        p1.textContent = jsonObj[i].text;
        p2.textContent = 'Date of create: ' + jsonObj[i].date_create;

        listItem.appendChild(h4);
        listItem.appendChild(p1);
        listItem.appendChild(p2);

        list.appendChild(listItem)
    }
    article.appendChild(list)
    articleSection.appendChild(article)
}