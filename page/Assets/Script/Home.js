const filterBtn = document.getElementById("Filters");

filterBtn.addEventListener("click", function() {
    filterBtn.classList.add("deployed");
});


document.addEventListener("DOMContentLoaded", function() {
    const navLinks = document.querySelectorAll("#nav a");
    const sections = Array.from(navLinks).map(link => document.querySelector(link.getAttribute("href")));

    function setActiveLink() {
        let currentActiveIndex = null;

        sections.forEach((section, index) => {
            const rect = section.getBoundingClientRect();

            if (rect.top <= 0 && rect.bottom >= 0) {
                currentActiveIndex = index;
            }
        });

        navLinks.forEach((link, index) => {
            if (index === currentActiveIndex) {
                link.classList.add("active");
            } else {
                link.classList.remove("active");
            }
        });
    }

    setActiveLink();

    window.addEventListener("scroll", () => {
        setActiveLink();
    });
});

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

    const artistDatesLink = card.querySelector(".datesButton");
    artistDatesLink.href = `/artistPage?artist=${artist.name}#ArtistDates`;

    const artistLink = card.querySelector(".artist-info");
    artistLink.href = `/artistPage?artist=${artist.name}`;


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

displayArtistCards(12).then(r => console.log(r));


// DATES


async function getUpcomingConcerts() {
    const artists = await fetchArtistData();
    const upcomingConcerts = [];

    artists.forEach((artist) => {
        for (const location in artist.datesLocations) {
            const dates = artist.datesLocations[location];
            dates.forEach((date) => {
                upcomingConcerts.push({
                    artistName: artist.name,
                    date: date,
                    location: location,
                });
            });
        }
    });

    return upcomingConcerts;
}


function sortConcertsByDate(concerts) {
    return concerts.sort((a, b) => {
        const [monthA, dayA, yearA] = a.date.split('-').map(Number);
        const [monthB, dayB, yearB] = b.date.split('-').map(Number);

        const dateA = new Date(yearA, monthA - 1, dayA);
        const dateB = new Date(yearB, monthB - 1, dayB);

        return dateB - dateA;
    });
}

async function createHomeDateCard(artistName, date, location) {
    const template = document.querySelector(".Date-card-template");
    const card = template.content.cloneNode(true);

    const artistElement = card.querySelector(".card-artist-name");
    artistElement.textContent = artistName;

    const dateElement = card.querySelector(".card-concert-date");
    dateElement.textContent = date;

    const locationElement = card.querySelector(".card-concert-location");
    const [city, country] = location.split("-");
    locationElement.textContent = city.replace("_", " ");

    const cityImage = card.querySelector(".DateCardImage");
    cityImage.src = await fetchPhotoUrl(location);

    const cardLink = card.querySelector(".home-Date-Card-link");
    cardLink.href = `https://www.google.fr/maps/place/${location}`;

    return card;
}


async function displayRecentConcerts() {
    const container = document.querySelector(".dates-container");
    const concerts = await getUpcomingConcerts();
    const sortedConcerts = sortConcertsByDate(concerts);
    const recentConcerts = sortedConcerts.slice(0, 10);

    for (const concert of recentConcerts) {
        const card = await createHomeDateCard(concert.artistName, concert.date, concert.location);
        container.appendChild(card);
    }
}

async function fetchPhotoUrl(searchTerm) {
    const apiUrl = `https://api.unsplash.com/search/photos?query=${searchTerm}-city&client_id=qeN1F7bV473dm1aW_F5u6nfnc-6BlCIfoeaTm8fSSBY`;

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

displayRecentConcerts();

// SEARCH BAR

function submitForm() {
    document.getElementById("searchForm").submit();
}
function UpdateHeaderLoginStatus() {
    const usernameCookie = document.cookie
        .split('; ')
        .find(row => row.startsWith('username='))
        ?.split('=')[1];

    if (usernameCookie) {
        document.querySelector('.btnLogin-popup').style.display = 'none';
        document.querySelector('#user-info').style.display = 'flex';
        document.querySelector('#username').textContent = usernameCookie;
    }
}

document.addEventListener('DOMContentLoaded', UpdateHeaderLoginStatus);

// document.getElementById('login-form').addEventListener('submit', (e) => {
//     e.preventDefault(); // Empêcher la soumission automatique du formulaire


// });

window.addEventListener('beforeunload', () => {
    document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
});

