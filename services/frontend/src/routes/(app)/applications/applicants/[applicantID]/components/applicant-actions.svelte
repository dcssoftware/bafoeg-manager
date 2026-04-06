<script>
  import { Button } from "$lib/components/Button";
</script>

<h2>Antragsteller Aktionen</h2>
<div class="actions-container">
  <div>
    <label>
      <span>Unterbinde alle automatischen Zahlungen</span>
      <input type="checkbox" />
    </label>
  </div>
  <div>
    <label>
      <span>Unterbinde den Eingang von weiteren BAföG Anträge</span>
      <input type="checkbox" />
    </label>
  </div>
  <div>
    <label>
      <span>Manuellen Rückzahlungsbetrag eintragen</span>
      <div>
        <input type="number" placeholder="Betrag in Euro" />
        <Button>Speichern</Button>
      </div>
    </label>
  </div>
  <div>
    <label>
      <span>Datenauskunft gemäß §15 DSGVO</span>
      <Button>Download</Button>
    </label>
  </div>
  <div>
    <label>
      <span>Datenlöschung gemäß §17 DSGVO</span>
      <Button>Delete</Button>
    </label>
  </div>
</div>

<style lang="sass">
  .actions-container
    display: grid
    grid-template-columns: repeat(2, 1fr)
    gap: 1rem
</style>
