fetch(`/teas`)
    .then(response => response.json())
    .then(data => {
        buildList(data);
    });


function buildList(teas) {
    if (!teas) {
        return;
    }

    teas.map((tea) => {
        const listElem = document.getElementById('tea-list')
        const liElem = document.createElement('li');

        const linkElem = document.createElement('a');
        linkElem.href = `/${tea.id}`
        linkElem.innerText = tea.id;
        linkElem.className = 'tea-link';

        const formElem = document.createElement('form');
        formElem.setAttribute('onsubmit', 'doPUT(event)')
        formElem.className = 'admin-form'

        const hiddenInput = document.createElement('input');
        hiddenInput.setAttribute('type', 'hidden');
        hiddenInput.setAttribute('name', 'teaId');
        hiddenInput.setAttribute('value', tea.id);

        const nameInput = document.createElement('input');
        nameInput.setAttribute('type',"text");
        nameInput.setAttribute('name',"teaName");
        nameInput.setAttribute('value',`${tea.teaName}`);
        nameInput.setAttribute('placeholder','Tea Name');

        const typeInput = document.createElement('input');
        typeInput.setAttribute('type',"text");
        typeInput.setAttribute('name',"teaType");
        typeInput.setAttribute('value',`${tea.teaType}`);
        typeInput.setAttribute('placeholder','Tea Type');

        const sizeInput = document.createElement('input');
        sizeInput.setAttribute('type',"text");
        sizeInput.setAttribute('name',"size");
        sizeInput.setAttribute('value',`${tea.size}`);
        sizeInput.setAttribute('placeholder','Tin Size');

        const colorInput = document.createElement('input');
        colorInput.setAttribute('type',"text");
        colorInput.setAttribute('name',"color");
        colorInput.setAttribute('value',`${tea.color}`);
        colorInput.setAttribute('placeholder','Tin Color');

        const inUseInput = document.createElement('input');
        inUseInput.setAttribute('type',"text");
        inUseInput.setAttribute('name',"inUse");
        inUseInput.setAttribute('value',`${tea.inUse}`);
        inUseInput.setAttribute('placeholder','In Use? (1, 0)');

        const temperatureInput = document.createElement('input');
        temperatureInput.setAttribute('type',"number");
        temperatureInput.setAttribute('name',"temperature");
        temperatureInput.setAttribute('value',`${tea.temperature}`);
        temperatureInput.setAttribute('placeholder','Tea Temperature');

        const portionWeightInput = document.createElement('input');
        portionWeightInput.setAttribute('type',"number");
        portionWeightInput.setAttribute('name',"portionWeight");
        portionWeightInput.setAttribute('value', `${tea.portionWeight}`);
        portionWeightInput.setAttribute('placeholder', 'Tea Portion Weight');

        const containerWeightInput = document.createElement('input');
        containerWeightInput.setAttribute('type',"number");
        containerWeightInput.setAttribute('name',"containerWeight");
        containerWeightInput.setAttribute('value', `${tea.containerWeight}`);
        containerWeightInput.setAttribute('placeholder', 'Tea Container Weight');

        const initialWeightInput = document.createElement('input');
        initialWeightInput.setAttribute('type',"number");
        initialWeightInput.setAttribute('name',"initialWeight");
        initialWeightInput.setAttribute('value', `${tea.initialWeight}`);
        initialWeightInput.setAttribute('placeholder', 'Tea Initial Weight');

        const brewingDurationInput = document.createElement('input');
        brewingDurationInput.setAttribute('type',"number");
        brewingDurationInput.setAttribute('name',"brewingDuration");
        brewingDurationInput.setAttribute('value', `${tea.brewingDuration}`);
        brewingDurationInput.setAttribute('placeholder', 'Brewing Duration');

        const shopNameInput = document.createElement('input');
        shopNameInput.setAttribute('type',"text");
        shopNameInput.setAttribute('name',"shopName");
        shopNameInput.setAttribute('value', `${tea.origin.shopName}`);
        shopNameInput.setAttribute('placeholder', 'Shop Name');

        const shopLocationInput = document.createElement('input');
        shopLocationInput.setAttribute('type',"text");
        shopLocationInput.setAttribute('name',"shopLocation");
        shopLocationInput.setAttribute('value', `${tea.origin.shopLocation}`);
        shopLocationInput.setAttribute('placeholder', 'Shop Location');

        const descriptionInput = document.createElement('textarea');
        descriptionInput.setAttribute('name', 'blendDescription');
        descriptionInput.innerText = tea.blendDescription;
        descriptionInput.setAttribute('placeholder', 'Blend Description');

        const submitInput = document.createElement("input");
        submitInput.setAttribute('type',"submit");
        submitInput.setAttribute('value',"Submit");

        const deleteButton = document.createElement('button');
        deleteButton.innerText = 'Delete Tea';
        deleteButton.setAttribute('onclick', `doDelete(event, '${tea.id}')`);
        deleteButton.setAttribute('type', 'button');

        formElem.append(
            nameInput,
            typeInput,
            temperatureInput,
            portionWeightInput,
            containerWeightInput,
            initialWeightInput,
            brewingDurationInput,
            shopNameInput,
            shopLocationInput,
            sizeInput,
            colorInput,
            inUseInput,
            descriptionInput,
            submitInput,
            deleteButton,
            hiddenInput,
        );

        const qrCodeElement = document.createElement('div');
        qrCodeElement.className = 'qrcode';
        new QRCode(qrCodeElement, {
            text: `${window.location.origin}/${tea.id}`,
            width: 128,
            height: 128,
            colorDark : "#000000",
            colorLight : "#ffffff",
            correctLevel : QRCode.CorrectLevel.H
        });

        liElem.append(
            linkElem,
            formElem,
            qrCodeElement
        );

        listElem.appendChild(liElem);

    })
}

function doDelete(event, id) {
    event.preventDefault();
    const requestOptions = {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' }
    };
    fetch(`/tea/${id}`, requestOptions)
        .then(response => response.json())
        .then(data => console.log(data) );
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
            "temperature": parseInt(event.target.elements.temperature.value),
            "portionWeight": parseInt(event.target.elements.portionWeight.value),
            "containerWeight": parseInt(event.target.elements.containerWeight.value),
            "initialWeight": parseInt(event.target.elements.initialWeight.value),
            "brewingDuration": parseInt(event.target.elements.brewingDuration.value),
            "teaName": event.target.elements.teaName.value,
            "teaType": event.target.elements.teaType.value,
            "color": event.target.elements.color.value,
            "size": event.target.elements.size.value,
            "inUse": parseInt(event.target.elements.inUse.value),
            "blendDescription": event.target.elements.blendDescription.value
        })
    };
    fetch(`/tea/${event.target.elements.teaId.value}`, requestOptions)
        .then(response => response.json())
        .then(data => console.log(data) );
}

function createNewTea(event) {
    event.preventDefault();
    const postForm = document.getElementById('postForm');
    const formedData = new FormData(postForm)
    let postObj = {};
    let origin = {};

    for (const [key, value] of formedData) {
        if(key === 'shopName' || key === 'shopLocation') {
            origin[key] = value;
            postObj.origin = origin;
        }
        if (key === 'temperature' || key === 'containerWeight' || key === 'initialWeight' || key === 'portionWeight' || key === 'brewingDuration' || key === 'inUse') {
            postObj[key] = parseInt(value);
        } else {
            postObj[key] = value;
        }
    }

    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(postObj)
    };

    fetch('/tea', requestOptions)
        .then(response => response.json())
        .then(data => console.log(data) );
}