import type { NavigationItemsType } from "$lib/navigation";

export function changeNavigationActiveElement(
    url: URL,
    navigation: NavigationItemsType
): number {
    let activeElementIndex: number = 0;
    navigation.forEach((navElement, index) => {
        const webPath = url.pathname + "/";
        const webPathRegex = new RegExp(`${navElement.href}\/.*`);
        const match = webPathRegex.test(webPath);

        if (match) {
            activeElementIndex = index
            return
        }
    });

    return activeElementIndex
}