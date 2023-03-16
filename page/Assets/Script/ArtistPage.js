function getArtistNameFromUrl() {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get("artist");
}

async function fetchArtistsData() {
    const response = await fetch("http://localhost:3000/api");
    return await response.json();
}

async function fetchArtistData(artistName) {
    const artists = await fetchArtistsData();

    if (artists && artists.length > 0) {
        return artists.find(
            (artist) => artist.name.toLowerCase() === artistName.toLowerCase()
        );
    }
    return null;
}


(async () => {
    const artistName = getArtistNameFromUrl();

    if (artistName) {
        const artist = await fetchArtistData(artistName);
        displayArtistInfo(artist);
    } else {
        console.error("Aucun nom d'artiste trouvé dans l'URL.");
    }
})();

function displayArtistInfo(artist) {
    document.getElementById("ArtistName").textContent = artist.name;
    document.getElementById("artistImage").src = artist.image;
    displayDateCards(artist);
}



async function createDateCard(date, location) {
    const template = document.getElementById("Date-card-template");
    const dateCard = template.content.cloneNode(true).querySelector(".DateCard");

    const dateElement = dateCard.querySelector(".DateCardDate");
    dateElement.textContent = date;

    const locationElement = dateCard.querySelector(".DateCardLocation");
    locationElement.textContent = location;

    const cityImage = dateCard.querySelector(".DateCardImage");
    cityImage.src = await fetchPhotoUrl(location);

    return dateCard;
}


async function displayDateCards(artistData) {
    const container = document.querySelector(".Date-container");

    for (const location in artistData.datesLocations) {
        const dates = artistData.datesLocations[location];

        for (const date of dates) {
            const dateCard = await createDateCard(date, location);
            container.appendChild(dateCard);
        }
    }
}

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