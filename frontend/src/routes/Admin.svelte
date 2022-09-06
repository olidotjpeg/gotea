<main>
    <div class="wrapper">
        <header class="header">
            <h1>GoTea Admin 😎</h1>
        </header>

        <div class="container">
            <ul id="tea-list">
                {#each teas as tea}
                    <li>
                        <a class="tea-link" href="/#/tea/{tea.id}">{tea.id}</a>
                        <form class="admin-form">
                            <input type="hidden" name="teaId" value="{tea.id}">
                            <input type="text" name="teaName" value="{tea.teaName}" placeholder="Tea Name">
                            <input type="text" name="teaType" value="{tea.teaType}" placeholder="Tea Type">
                            <input type="text" name="size" value="{tea.size}" placeholder="Tin Size">
                            <input type="text" name="color" value="{tea.color}" placeholder="Tin Color">
                            <input type="text" name="inUse" value="{tea.inUse}" placeholder="In use? (1, 0)">
                            <input type="text" name="temperature" value="{tea.temperature}" placeholder="Brew temperature">
                            <input type="text" name="portionWeight" value={tea.portionWeight} placeholder="Weight per Portion">
                            <input type="text" name="containerWeight" value={tea.containerWeight} placeholder="Container Weight">
                            <input type="text" name="initialWeight" value={tea.initialWeight} placeholder="Initial Weight of Container">
                            <input type="text" name="brewingDuration" value={tea.brewingDuration} placeholder="Brewing Duration">
                            <input type="text" name="shopLocation" value={tea.origin.shopLocation} placeholder="Shop Location">
                            <input type="text" name="shopName" value={tea.origin.shopName} placeholder="Shop Name">
                            <textarea  name="blendDescription" value={tea.blendDescription} rows="3" placeholder="Blend Description"></textarea>

                            <input type="submit" value="Submit">
                            <button type="button">Delete</button>
                        </form>
                        {tea.teaName}
                    </li>
                {/each}
            </ul>
        </div>

        <!-- <div class="container">
            <h4>Add new Teas here</h4>
            <form onsubmit="createNewTea(event)" class="admin-form" id="postForm">
                <input name="teaName" placeholder="Tea Name" type="text" />
                <input name="shopName" placeholder="Shop Name" type="text" />
                <input name="shopLocation" placeholder="Shop Location" type="text" />
                <input name="temperature" placeholder="Temperature" type="text" />
                <input name="portionWeight" placeholder="Portion Weight" type="text" />
                <input name="containerWeight" placeholder="Container Weight" type="text" />
                <input name="initialWeight" placeholder="Initial Weight" type="text" />
                <input name="brewingDuration" placeholder="Brewing Duration (In Seconds)" type="text" />
                <input name="size" placeholder="Size of the tin (Large, Small)" type="text" />
                <input name="color" placeholder="Color of the tin" type="text" />
                <input name="inUse" placeholder="In Use? (1, 0)" type="text" />
                <input type="submit" />
            </form>
        </div> -->
    </div>
</main>

<script lang="ts">
    import type {Tea} from '../interfaces';

    let teas: Tea[] = [];

    fetch('http://localhost:8000/teas').then((res) => res.json()).then((data) => teas = data);
</script>

<style>
    .header {
        background: purple;
        color: #fff;
        height: 75px;
        margin-bottom: 24px;
        display: flex;
        align-items: center;
        padding-left: 24px;
    }

    .container {
        max-width: 310px;
        margin: 0 auto;
    }

    #tea-list li {
        list-style-type: none;
    }

    .admin-form {
        display: flex;
        flex-direction: column;
        max-width: 310px;
    }

    .tea-link {
        display: block;
        margin-bottom: 50px
    }

    .qrcode {
        margin-bottom: 20px;
    }

    .admin-form input {
        margin-bottom: 12px;
        border-radius: 5px;
        background: #fafafa;
        border: 1px solid #ccc;
        padding: 4px;
    }

    .admin-form input[type="submit"] {
        background: #36BFB1;
        border: none;
        color: #fff;
        font-weight: bold;
        border-radius: 3px;
        font-size: 18px;
        cursor: pointer;
    }
</style>