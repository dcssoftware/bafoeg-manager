<script lang="ts">
  import Chart from "chart.js/auto";

  let chartElement: HTMLCanvasElement | null;
  let chartElement2: HTMLCanvasElement | null;

  $effect(() => {
    if (chartElement != null) {
      new Chart(chartElement, {
        type: "line",
        data: {
          labels: ["Red", "Blue", "Yellow", "Green", "Purple", "Orange"],
          datasets: [
            {
              label: "Bearbeitete Anträge",
              data: [30285, 15712, 26132, 2001, 9555, 38005],
              borderWidth: 1,
            },
            {
              label: "Eingehende Anträge",
              data: [15000, 12000, 16500, 3000, 10000, 19000],
              borderWidth: 1,
            },
          ],
        },
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    }
    if (chartElement2 != null) {
      new Chart(chartElement2, {
        type: "bar",
        data: {
          labels: [
            "Oktober",
            "November",
            "Dezember",
            "Januar",
            "Februar",
            "März",
          ],
          datasets: [
            {
              label: "Gesamtauszahlungen in Euro",
              data: [5000000, 6000000, 7000000, 4000000, 3500000, 7000000],
              borderWidth: 1,
            },
          ],
        },
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    }
  });
</script>

<div class="main-charts">
  <div class="application-history">
    <canvas class="applications-history-chart" bind:this={chartElement}
    ></canvas>
    <div>
      <h2>Antragshistorie</h2>
      <p>
        Hier sehen wir das Verhältnis den Daten der gestellten Anträge und den
        der bearbeiteten Anträge
      </p>
    </div>
  </div>
  <div class="payment-history">
    <canvas class="payment-history-chart" bind:this={chartElement2}></canvas>
    <div>
      <h2>Zahlungshistorie</h2>
      <p>
        Hier sehen wir das ausgegebene Geld im Verhältnis der letzten Monate
      </p>
    </div>
  </div>
</div>

<style lang="sass">
  .main-charts
    display: flex
    flex-direction: row
    flex-wrap: wrap
    gap: 2rem
    &>div
      max-width: 30%
</style>
