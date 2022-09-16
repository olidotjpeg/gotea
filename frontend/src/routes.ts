import Admin from './routes/Admin.svelte'
import Details from './routes/Details.svelte'
import NotFound from './routes/NotFound.svelte'

export const routes = {
    // Exact path
    '/admin': Admin,

    // Using named parameters, with last being optional
    '/tea/:teaId': Details,

    // Catch-all
    // This is optional, but if present it must be the last
    '*': NotFound,
}