import { IconBook, IconEnvelope, IconHome, IconPieChart, IconSchool, IconStudent, IconCartFlatbed, IconUsers, IconSackDollars, IconSitemap } from "$lib/components/Icons";
import { _ } from "svelte-i18n";

export type NavigationItemsType = NavigationItemType[]

export interface NavigationItemType {
    name: string;
    href: string;
    icon: any;
    permission?: string;
    subelements?: {
        name: string;
        href: string;
        permission?: string;
    }[];
}

export function getNavigationItems(store: CallableFunction): NavigationItemsType {
    return [
        {
            // name: "Startseite",
            name: store("page.navigation.home"),
            href: "/",
            icon: IconHome,
        },
        {
            // name: "Anträge",
            name: store("page.navigation.application.header"),
            href: "/applications",
            icon: IconStudent,
            subelements: [
                {
                    name: store("page.navigation.application.labels"),
                    href: "/applications/labels",
                },
                {
                    name: store("page.navigation.applicants.management"),
                    href: "/applications/applicants",
                },
            ],
        },
        // {
        //     name: store("page.navigation.auto-aprovals.header"),
        //     href: "/auto-approvals",
        //     icon: IconPeopleApproval,
        // },
        {
            name: store("page.navigation.educational-institutions.header"),
            href: "/school-management",
            icon: IconSchool,
        },
        {
            // name: "Anträge",
            name: store("page.navigation.application-rag.header"),
            href: "/rag",
            icon: IconBook,
            subelements: [
                {
                    name: store("page.navigation.application-rag.studierendenbafoeg"),
                    href: "/rag/studierenden-bafoeg-regeln",
                    permission: "read:rag-management-studierenden-files",
                },
                {
                    name: store("page.navigation.application-rag.schuelerbafoeg"),
                    href: "/rag/schueler-bafoeg-regeln",
                    permission: "read:rag-management-schueler-files",
                },
            ],
        },
        // {
        //     name: store("page.navigation.questions-mailbox.header"),
        //     href: "/inquiries-mailbox",
        //     icon: IconEnvelope,
        // },
        {
            name: store("page.navigation.e-akte.header"),
            href: "/eakte",
            icon: IconCartFlatbed,
        },
        // {
        //     name: store("page.navigation.finances.header"),
        //     href: "/finances",
        //     icon: IconSackDollars,
        // },
        // {
        //     name: store("page.navigation.statistics.header"),
        //     href: "/statistics",
        //     icon: IconPieChart,
        // },
        {
            name: store("page.navigation.organization.header"),
            href: "/organization",
            icon: IconSitemap,
        },
        {
            name: store("page.navigation.user-management.header"),
            href: "/user-management",
            permission: "read:user-management",
            icon: IconUsers,
            // subelements: [
            //     {
            //         name: "Rollen Verwaltung",
            //         href: "/user-management/roles",
            //         permission: "read:user-management",
            //     }
            // ]
        },
    ]
};