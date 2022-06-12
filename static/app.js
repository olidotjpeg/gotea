let expanded = false;

function loadData() {
    const teaLocation = window.location.pathname;
    fetch(`/tea${teaLocation}`)
        .then(response => response.json())
        .then(data => {
            document.getElementById('teaType').innerText = data.teaType;
            document.getElementById('originShop').innerText = data.origin.shopName;
            document.getElementById('teaName').innerText = data.teaName;
            document.getElementById('temperature').innerText = `${data.temperature}°C`;
            document.getElementById('brewTime').innerText = `${data.brewingDuration}min`;
            document.getElementById('portionWeight').innerText = `${data.portionWeight}g/l`;
        });
}

function expandMenu() {
    const expandElement = document.querySelector('.expand-element');

    expanded = !expanded;
    if (expanded) {
        expandElement.classList.add('expanded');
    } else {
        expandElement.classList.remove('expanded');
    }
}

loadData();

function debugTeas() {
    fetch(`http://localhost:8000/teas`)
        .then(response => response.json())
        .then(data => console.log(data));
}