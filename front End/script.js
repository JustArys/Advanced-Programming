const apiUrl = 'http://127.0.0.1:8080/books/'; // Replace with your actual Go backend URL

function getAllBooks() {
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            console.log('Response:', data); // Log the response
            displayBooks(data.data);
        })
        .catch(error => console.error('Error:', error));
}


function displayBooks(data) {
    const booksList = document.getElementById('books-list');
    booksList.innerHTML = '';

    if (Array.isArray(data)) {
        data.forEach(book => {
            const li = document.createElement('li');
            li.textContent = `${book.title} by ${book.author} (${book.year})`;
            booksList.appendChild(li);
        });
    } else {
        // If the response is not an array, you can handle it as needed
        console.error('Invalid response format for displayBooks:', data);
    }
}

// Other functions for addBook, updateBook, and deleteBook can remain the same.


// Other functions for updateBook and deleteBook can remain the same.

function addBook(event) {
    event.preventDefault();

    const title = document.getElementById('title').value;
    const author = document.getElementById('author').value;
    const year = document.getElementById('year').value;

    const newBook = {
        title: title,
        author: author,
        year: year,
    };

    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newBook),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Success:', data);
            getAllBooks(); // Refresh the books list after adding a new book
        })
        .catch(error => console.error('Error:', error));
}

// Other functions for updateBook and deleteBook can be similar with appropriate modifications.
