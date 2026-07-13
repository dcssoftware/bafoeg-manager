<script lang="ts">
    import { AlluvialChart } from '@carbon/charts-svelte'
    import '@carbon/charts-svelte/styles.css'
    const options = {
        alluvial: {
            nodes: [
                {
                    name: 'Antragsteller',
                    category: 'Antragsteller'
                },
                {
                    name: 'München',
                    category: 'Studierendenwerk'
                },
                {
                    name: 'TUM München 1. Semester',
                    category: 'Ausbildungsstätte'
                },
                {
                    name: 'Beispiel Gymnasium, Klasse 12',
                    category: 'Ausbildungsstätte'
                },
            ]
        },
        height: '600px',
        color: {
            scale: {
                "Antragsteller": "var(--color-green)",
                "München": "var(--color-green)",
                "TUM München": "var(--color-green)",
                "Beispiel Gymnasium": "var(--color-green)",
            },
            gradient: {
                enabled: false
            }
        }
    }
    const data = [
        {
            source: "Antragsteller",
            target: "München",
            value: 1000
        },
        {
            source: "München",
            target: "TUM München 1. Semester",
            value: 300
        },
        {
            source: "München",
            target: "Beispiel Gymnasium, Klasse 12",
            value: 700
        }
    ]
</script>

<div>
    <AlluvialChart {options} {data} />
</div>