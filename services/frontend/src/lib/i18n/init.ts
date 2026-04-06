import { addMessages, init } from "svelte-i18n";
import { allowedLanguages, getLocale } from ".";

import { DElanguage, ENlanguage } from "$lib/i18n/";

export function initI18n() {

    // configure i18n
    addMessages("en", ENlanguage);
    addMessages("de", DElanguage);
    init({
        fallbackLocale: "en",
        initialLocale: getLocale(allowedLanguages),
    });
}