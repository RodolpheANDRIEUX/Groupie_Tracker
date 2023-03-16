async function fetchArtistData() {
    const response = await fetch("http://localhost:3000/api");
    return await response.json();
}

function createArtistCard(artist) {
    const template = document.getElementById("artist-card-template");
    const card = template.content.cloneNode(true);

    const artistImage = card.querySelector(".artist-image");
    artistImage.src = artist.image;
    artistImage.alt = `Image of ${artist.name}`;

    const artistName = card.querySelector(".artist-name");
    artistName.textContent = artist.name;

    const artistDate = card.querySelector(".artist-date");
    artistDate.textContent = artist.creationDate;

    return card;
}

function sortArtistsByName(artists) {
    return artists.sort((a, b) => {
        if (a.name.toLowerCase() < b.name.toLowerCase()) {
            return -1;
        }
        if (a.name.toLowerCase() > b.name.toLowerCase()) {
            return 1;
        }
        return 0;
    });
}

function sortArtistsByDate(artists) {
    return artists.sort((a, b) => {
        if (a.creationDate < b.creationDate) {
            return -1;
        }
        if (a.creationDate > b.creationDate) {
            return 1;
        }
        return 0;
    });
}

async function displayArtistCards(limit) {
    const container = document.querySelector(".grid-container");
    const artists = await fetchArtistData();

    if (artists && artists.length > 0) {
        const sortedArtists = sortArtistsByName(artists);

        sortedArtists.slice(0, limit).forEach((artist) => {
            const artistCard = createArtistCard(artist);
            container.appendChild(artistCard);
        });
    }
}

document.getElementById("More").addEventListener("click", async () => {
    const container = document.querySelector(".grid-container");
    container.innerHTML = "";
    await displayArtistCards(Infinity);
    document.getElementById("More").style.display = "none";
});

async function fetchPhotoUrl(searchTerm) {
    const apiUrl = `https://api.unsplash.com/search/photos?query=${searchTerm}&client_id=qeN1F7bV473dm1aW_F5u6nfnc-6BlCIfoeaTm8fSSBY`;

    try {
        const response = await fetch(apiUrl);
        const data = await response.json();

        if (data.results && data.results.length > 0) {
            const firstPhoto = data.results[0];
            return firstPhoto.urls.raw;
        } else {
            console.log("Aucune photo trouvée pour ce terme de recherche.");
            return null;
        }
    } catch (error) {
        console.error("Erreur lors de la récupération des données de l'API Unsplash :", error);
        return null;
    }
}

displayArtistCards(12).then(r => console.log(r));