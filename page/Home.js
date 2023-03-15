async function fetchArtistData() {
    const response = await fetch("https://groupietrackers.herokuapp.com/api/artists");
    return await response.json();
}

function createArtistCard(artist) {
    const card = document.createElement("div");
    card.className = "card";

    const artistName = document.createElement("h3");
    artistName.textContent = artist.name;
    card.appendChild(artistName);

    // ect (images et tout)

    return card;
}

async function displayArtistCards() {
    const container = document.querySelector(".container");
    const artists = await fetchArtistData();

    if (artists && artists.length > 0) {
        artists.forEach((artist) => {
            const artistCard = createArtistCard(artist);
            container.appendChild(artistCard);
        });
    }
}

displayArtistCards();