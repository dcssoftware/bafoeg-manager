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
                    name: 'TUM München',
                    category: 'Ausbildungsstätte'
                },
                {
                    name: 'Beispiel Gymnasium',
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
                enabled: true
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
            target: "TUM München",
            value: 300
        },
        {
            source: "München",
            target: "Beispiel Gymnasium",
            value: 700
        }
    ]
</script>

<div>
    <AlluvialChart {options} {data} />
</div>