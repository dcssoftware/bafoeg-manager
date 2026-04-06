<script>
  import { Button } from "$lib/components/Button";
  import { PieChart } from "$lib/components/Chart-Pie";
</script>

<div class="header">
  <h1>Statistiken</h1>
  <Button>Ausdrucken</Button>
</div>

<PieChart></PieChart>

<div class="disclaimer">
  <p>
    ⚠️ Die Statistiken sind nur grob zur Demonstration und enthalten keine
    echten Daten! ⚠️
  </p>
</div>

<style lang="sass">
  .header
    display: flex
    justify-content: space-between
    align-items: center
    margin-bottom: 2rem

  .disclaimer
    text-align: center
    margin-top: 2rem
    p
      font-size: 2rem
</style>
