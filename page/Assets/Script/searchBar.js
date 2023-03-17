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
            const filteredResults = [];

            artists.forEach(artist => {
                if (artist.name.toLowerCase().includes(searchTerm)) {
                    filteredResults.push({ name: artist.name, artist: artist.name });
                } else {
                    artist.members.forEach(member => {
                        if (member.toLowerCase().includes(searchTerm)) {
                            filteredResults.push({ name: member, artist: artist.name });
                        }
                    });
                }
            });

            if (filteredResults.length === 0) {
                results.style.display = 'none';
            } else {
                results.style.display = 'block';
                filteredResults.forEach(result => {
                    const link = document.createElement('a');
                    link.textContent = result.name;
                    link.href = `http://localhost:3000/artistPage?artist=${result.artist}`;
                    results.appendChild(link);
                });
            }
        });
    }
});