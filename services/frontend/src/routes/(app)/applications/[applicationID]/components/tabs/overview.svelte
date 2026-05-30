<script lang="ts">
  import { FinancialTable } from "$lib/components/FinancialTable";

  interface Props {
    gotoAiReport: () => void;
  }
  let { gotoAiReport }: Props = $props();

  function elternAnrechenbar(value: number): number {
    let anrechenbar = value / 2;
    if (anrechenbar < 0) {
      return 0;
    }
    return -anrechenbar;
  }
</script>

<h2>Übersicht</h2>

<div class="alert-warning">
  <strong>Hinweis:</strong> Die hier dargestellten Berechnungen dienen nur zu Illustrationszwecken
  und spiegeln (noch) nicht die tatsächlichen Ergebnisse wider. Die Berechnung wird
  in späteren Versionen korrekt dargestellt werden.
</div>

<div class="overview-container">
  <!-- <div class="ai-precheck">
    <h3>KI Auswertung</h3>
    <div class="metrics-grid">
      <div class="metric-card critical">
        <span class="label">Kritische Fehler</span>
        <span class="number">2</span>
      </div>
      <div class="metric-card warning">
        <span class="label">Warnungen</span>
        <span class="number">5</span>
      </div>
      <div class="metric-card precheck">
        <span class="label">System Score</span>
        <span class="number">87%</span>
      </div>
    </div>
    <div class="goto-ai-report">
      <a onclick={gotoAiReport}>
        <span>Den vollen KI-Report ansehen</span>
      </a>
    </div>
  </div> -->
  <div class="bedarfsanalyse">
    <h3>Bedarfsanalyse</h3>
    <div class="calculations">
      <FinancialTable
        table={{
          title: "Bedarfssatz",
          values: [
            { header: "Grundbedarf", value: 449.0 },
            { header: "Zuschlag Krankenversicherung", value: 122.0 },
            { header: "Nicht bei den Eltern wohnend", value: 380.0 },
          ],
          ergebnis_title: "Bedarfssatz Gesamt",
        }}
      />
      <FinancialTable
        table={{
          title: "Eltern Finanzen",
          values: [
            { header: "Mutter Einkommen", value: 3000.0 },
            { header: "Mutter Freibetrag", value: -2400.0 },
            { header: "Vater Einkommen", value: 0.0 },
            { header: "Vater Freibetrag", value: 0.0 },
            { header: "Geschwister Freibetrag", value: -730.0 },
          ],
          ergebnis_title: "Eltern Finanzen Gesamt",
        }}
      />
      <FinancialTable
        table={{
          title: "Fördergelder",
          values: [
            { header: "Kindergeld", value: 200 },
            { header: "Halbwaisenrente", value: 120 },
          ],
        }}
      />
    </div>
    <br /><br />
    <FinancialTable
      table={{
        title: "Ergebnis",
        values: [
          { header: "Eigene Finanzen", value: 951 },
          { header: "Eltern Finanzen", value: elternAnrechenbar(-130) },
        ],
      }}
    />
  </div>
</div>

<style lang="sass">
.alert-warning
    background-color: var(--color-red-80)
    color: var(--font-color-white)
    display: flex
    flex-direction: column
    gap: 1rem
    padding: 2rem
    border-radius: 5px
    font-size: 1rem

    strong
      font-size: 1.5rem
      font-weight: bold

.overview-container
  .ai-precheck
    .goto-ai-report
      display: flex
      justify-content: flex-end
      margin-top: 1rem
      a
        display: block
        padding: 0.5rem 1rem
        background-color: var(--background-color-secondary)
        cursor: pointer
  .metrics-grid
    display: grid
    grid-template-columns: repeat(3, 1fr)
    gap: 1rem
  
    .metric-card
      background-color: var(--background-color-secondary);
      padding: 1rem 1rem
      border-radius: 5px
      display: flex
      justify-content: space-between
      align-items: center
      font-size: 1.2rem
      border-bottom: 3px solid

      .number
        width: 4rem
        height: 3rem
        text-align: center
        align-content: center
        border-radius: 5px
      &.critical
        border-color: var(--color-red)
        .number
          background-color: var(--color-red)
          color: var(--font-color-white)
      &.warning
        border-color: var(--color-yellow)
        .number
          background-color: var(--color-yellow)
          color: var(--font-color-black)
      &.precheck
        border-color: var(--color-blue)
        .number
          background-color: var(--color-blue)
          color: var(--font-color-white)

  .bedarfsanalyse
    .calculations
      display: grid
      grid-template-columns: repeat(3, 1fr)
      gap: 1rem
      .own-financials, .parents-financials, .goverment-funds
        background-color: var(--background-color-secondary)
        padding: 1rem
        border-radius: 5px
        display: flex
        flex-direction: column
        table
          width: 100%
          margin-top: auto
          th, td
            padding: .5rem 0
            text-align: left
            border-bottom: 1px solid var(--font-color)
          td
            text-align: right
          tr:last-of-type
            th, td
              border-bottom: none



</style>
