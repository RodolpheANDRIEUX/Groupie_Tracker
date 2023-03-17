const search = document.getElementById('search');
const results = document.getElementById('results');

function fetchArtistData2() {
    return fetch('http://localhost:3000/api')
        .then(response => response.json());
}

search.addEventListener('input', (event) => {
    const searchTerm = event.target.value.toLowerCase();
    results.innerHTML = '';

    if (searchTerm === '') {
        results.style.display = 'none';
    } else {
        fetchArtistData2().then(artists => {
            const filteredArtists = artists.filter(artist => artist.name.toLowerCase().includes(searchTerm));

            if (filteredArtists.length === 0) {
                results.style.display = 'none';
            } else {
                results.style.display = 'block';
                filteredArtists.forEach(artist => {
                    const link = document.createElement('a');
                    link.textContent = artist.name;
                    link.href = `http://localhost:3000/artistPage?artist=${artist.name}`;
                    results.appendChild(link);
                });
            }
        });
    }
});
