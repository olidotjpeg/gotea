<main>
    <div class="wrapper">
        <div class="tea-card">
            {#if tea}
                <h1 id="teaName">{tea.teaName}</h1>
                <div class="tea-meta">
                    <p id="teaType">{tea.teaType}</p>
                    <p id="originShop">{tea.origin.shopName}</p>
                </div>
                <img class="tea-picture" src="tea-picture.png" alt="This is a tea">
                <ul class="tea-guide">
                    <li>
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M8 5a4 4 0 1 1 8 0v5.255a7 7 0 1 1-8 0V5zm1.144 6.895a5 5 0 1 0 5.712 0L14 11.298V5a2 2 0 1 0-4 0v6.298l-.856.597zm1.856.231V5h2v7.126A4.002 4.002 0 0 1 12 20a4 4 0 0 1-1-7.874zM12 18a2 2 0 1 0 0-4 2 2 0 0 0 0 4z"/></svg>
                        <span id="temperature">{tea.temperature}</span>
                    </li>
                    <li>
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M17.618 5.968l1.453-1.453 1.414 1.414-1.453 1.453a9 9 0 1 1-1.414-1.414zM12 20a7 7 0 1 0 0-14 7 7 0 0 0 0 14zM11 8h2v6h-2V8zM8 1h8v2H8V1z"/></svg>
                        <span id="brewTime">{tea.brewingDuration}</span>
                    </li>
                    <li>
                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0H24V24H0z"/><path d="M6 2c0 .513.49 1 1 1h10c.513 0 1-.49 1-1h2c0 1.657-1.343 3-3 3h-4l.001 2.062C16.947 7.555 20 10.921 20 15v6c0 .552-.448 1-1 1H5c-.552 0-1-.448-1-1v-6c0-4.08 3.054-7.446 7-7.938V5H7C5.34 5 4 3.66 4 2h2zm6 7c-3.238 0-6 2.76-6 6v5h12v-5c0-3.238-2.762-6-6-6zm0 2c.742 0 1.436.202 2.032.554l-2.74 2.739c-.39.39-.39 1.024 0 1.414.361.36.929.388 1.32.083l.095-.083 2.74-2.739c.351.596.553 1.29.553 2.032 0 2.21-1.79 4-4 4s-4-1.79-4-4 1.79-4 4-4z"/></svg>
                        <span id="portionWeight">{tea.portionWeight}</span>
                    </li>
                </ul>
                <footer class="tea-footer">
                    <div class="footer-content">
                        {#if tea.link}
                            <a href="{tea.link}" class="link-to-tea">
                                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M18.364 15.536L16.95 14.12l1.414-1.414a5 5 0 1 0-7.071-7.071L9.879 7.05 8.464 5.636 9.88 4.222a7 7 0 0 1 9.9 9.9l-1.415 1.414zm-2.828 2.828l-1.415 1.414a7 7 0 0 1-9.9-9.9l1.415-1.414L7.05 9.88l-1.414 1.414a5 5 0 1 0 7.071 7.071l1.414-1.414 1.415 1.414zm-.708-10.607l1.415 1.415-7.071 7.07-1.415-1.414 7.071-7.07z"/></svg>
                            </a>
                        {/if}
                        <button on:click|preventDefault={expandMenu} class="expand-card">
                            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24"><path fill="none" d="M0 0h24v24H0z"/><path d="M12 13.172l4.95-4.95 1.414 1.414L12 16 5.636 9.636 7.05 8.222z"/></svg>
                        </button>
                    </div>
                    <div class="{expanded ? 'expanded' : ''} expand-element">
                        <span>{tea.blendDescription}</span>
                    </div>
                </footer>
            {/if}
        </div>
        <div class="background"></div>
    </div>
</main>

<script lang="ts">
    import type { Tea } from "../interfaces";

    export let params: {
        teaId: string;
    }

    let tea: Tea;
    let expanded = false;

    fetch(`http://localhost:8000/tea/${params.teaId}`).then((res) => res.json()).then((data) => {
        if (data) {
            tea = data
            console.log(tea);
        }
    });

    function expandMenu() {
        expanded = !expanded;
    }
</script>

<style>
.wrapper {
    width: 100vw;
    height: 100vh;
    position: relative;
    overflow: hidden;
}

.background {
    position: absolute;
    width: 110%;
    height: 110%;
    top: -20px;
    left: -20px;
    right: 0;
    bottom: 0;

    background-image: url("https://images.unsplash.com/photo-1561296180-e8100fd714e2?crop=entropy&cs=tinysrgb&fm=jpg&ixlib=rb-1.2.1&q=80&raw_url=true&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2734");
    filter: blur(10px);
}

.tea-card {
    width: 410px;
    position: absolute;
    z-index: 2;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);

    padding: 19px 16px;

    background: #FFFFFF;
    box-shadow: 0 1px 1px rgba(0, 0, 0, .14), 0 2px 1px rgba(0, 0, 0, .12), 0 1px 3px rgba(0, 0, 0, .2);
    border-radius: 4px;
}

.tea-card h1 {
    font-size: 24px;
    margin-bottom: 6px;
}

.tea-meta {
    font-size: 14px;
    display: flex;
    justify-content: space-between;
    color: rgba(0, 0, 0, .6)
}

.tea-picture {
    width: 100%;
}

.tea-guide {
    list-style-type: none;
    display: flex;
    justify-content: space-between;
    padding: 0;
}

.tea-guide li {
    display: flex;
    align-items: center;
}

.expand-element span {
    display: block;
    height: 0;
    opacity: 0;
    transition: height ease-in-out .3s, opacity ease-in-out .3s;
}

.expand-element.expanded span {
    height: 50px;
    opacity: 1;
}

.footer-content {
    margin-top: 8px;
    display: flex;
    justify-content: space-between;
}
</style>