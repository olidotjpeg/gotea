fetch(`/teas`)
    .then(response => response.json())
    .then(data => {
        buildList(data);
    });


function buildList(teas) {
    teas.map((tea) => {
        const listElem = document.getElementById('tea-list')
        const liElem = document.createElement('li');

        const linkElem = document.createElement('a');
        linkElem.href = `/${tea.id}`
        linkElem.innerText = tea.id;

        const formElem = document.createElement('form');
        formElem.setAttribute('onsubmit', 'doPUT(event)')

        const hiddenInput = document.createElement('input');
        hiddenInput.setAttribute('type', 'hidden');
        hiddenInput.setAttribute('name', 'teaId');
        hiddenInput.setAttribute('value', tea.id)

        const nameInput = document.createElement('input');
        nameInput.setAttribute('type',"text");
        nameInput.setAttribute('name',"teaName");
        nameInput.setAttribute('value',`${tea.teaName}`);

        const temperatureInput = document.createElement('input');
        temperatureInput.setAttribute('type',"text");
        temperatureInput.setAttribute('name',"temperature");
        temperatureInput.setAttribute('value',`${tea.temperature}`);

        const portionWeightInput = document.createElement('input');
        portionWeightInput.setAttribute('type',"text");
        portionWeightInput.setAttribute('name',"portionWeight");
        portionWeightInput.setAttribute('value', `${tea.portionWeight}`);

        const containerWeightInput = document.createElement('input');
        containerWeightInput.setAttribute('type',"text");
        containerWeightInput.setAttribute('name',"containerWeight");
        containerWeightInput.setAttribute('value', `${tea.containerWeight}`);

        const initialWeightInput = document.createElement('input');
        initialWeightInput.setAttribute('type',"text");
        initialWeightInput.setAttribute('name',"initialWeight");
        initialWeightInput.setAttribute('value', `${tea.initialWeight}`);

        const brewingDurationInput = document.createElement('input');
        brewingDurationInput.setAttribute('type',"text");
        brewingDurationInput.setAttribute('name',"brewingDuration");
        brewingDurationInput.setAttribute('value', `${tea.brewingDuration}`);

        const shopNameInput = document.createElement('input');
        shopNameInput.setAttribute('type',"text");
        shopNameInput.setAttribute('name',"shopName");
        shopNameInput.setAttribute('value', `${tea.origin.shopName}`);

        const shopLocationInput = document.createElement('input');
        shopLocationInput.setAttribute('type',"text");
        shopLocationInput.setAttribute('name',"shopLocation");
        shopLocationInput.setAttribute('value', `${tea.origin.shopLocation}`);

        const submitInput = document.createElement("input");
        submitInput.setAttribute('type',"submit");
        submitInput.setAttribute('value',"Submit");

        formElem.append(
            nameInput,
            hiddenInput,
            temperatureInput,
            portionWeightInput,
            containerWeightInput,
            initialWeightInput,
            brewingDurationInput,
            shopNameInput,
            shopLocationInput,
            submitInput
        );

        liElem.append(
            linkElem,
            formElem,
        );

        listElem.appendChild(liElem);
    })
}

function doPUT(event) {
    event.preventDefault();
    const requestOptions = {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            "origin": {
                "shopName": event.target.elements.shopName.value,
                "shopLocation": event.target.elements.shopLocation.value
            },
            "temperature": event.target.elements.temperature.value,
            "portionWeight": event.target.elements.portionWeight.value,
            "containerWeight": event.target.elements.containerWeight.value,
            "initialWeight": event.target.elements.initialWeight.value,
            "brewingDuration": event.target.elements.brewingDuration.value,
            "teaName": event.target.elements.teaName.value
        })
    };
    fetch(`/tea/${event.target.elements.teaId.value}`, requestOptions)
        .then(response => response.json())
        .then(data => console.log(data) );
}