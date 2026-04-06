<script lang="ts">
  import { ApplicationIndicator } from "$lib/components/ApplicationIndicator";
  import { ApplicationComments } from "$lib/components/ApplicationComments";
  import { FinancialTable } from "$lib/components/FinancialTable";

  let data = $state({
    noch_nich_bewilligte_foerderungen: ["AFBG"],
    vorraussichtliche_einnahmen: 123.45,
    bestehende_arbeitsverhaeltnisse_brutto: 0,
    bestehende_arbeitsverhaeltnisse_arbeitgeberanteil_vermoegenswirksame_leistungen: true,
    ausbildungs_und_praktikumsverguetung_brutto: 0,
    einkuenfte_aus_selbststaendiger_arbeit_gewerbe: 0,
    einnahmen_aus_kapitalvermoegen: 0,
    waisenrente_und_sonstige_renten: 0,
    ausbildungsbeihilfen_und_gleichartige_leistungen: 0,
    unterhaltsleistungen_nicht_der_eltern: 0,
    unterhalts_vorschussleistungen: 0,
    weitere_einnahmen: 0,
  });
</script>

<h2 class="header">Eigene Finanzen</h2>

<div class="form">
  <div class="calculations">
    <FinancialTable
      table={{
        title: "Bedarfsberechnung",
        values: [
          { header: "Einkommen Gesamt", value: 1523.89 },
          { header: "Unrelevante Abzüge (z.B. VWL)", value: 244.2 },
          {
            header: "Zu beachtendes Einkommen",
            value: 1340.14,
            isResult: true,
          },
          { header: "Sonstige Abgaben", value: 120 },
          { header: "Antragsteller Netto", value: 1120, isResult: true },
        ],
      }}
    />
  </div>
  <div class="own-income">
    <h3 class="header">Eigenes Einkommen</h3>
    <div class="list">
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={51} />
        </div>
        <div class="content">
          <div>
            Geförderte Altersvorsorgebeiträge nach §81 Einkommensgesetz (EStG)
            (z.B. Riester Renter)
          </div>
          <div class="user-value">
            <input
              type="text"
              value={data.noch_nich_bewilligte_foerderungen.join(", ")}
              disabled
            />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>

      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={52} />
        </div>
        <div class="content">
          <div>
            Im oben genannten Bewilligungszeitraum werde ich voraussichtlich
            Einnahmen erzielen
          </div>
          <div class="user-value">
            <input
              type="text"
              disabled
              class="number"
              value={data.vorraussichtliche_einnahmen}
            />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>

      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={53} />
        </div>
        <div class="content">
          <div>
            Bruttoeinnahmen aus bestehenden oder ruhenden Arbeitsverhältnissen,
            Gelegenheitsarbeiten, Ferien-, Minijobs
          </div>
          <div class="user-value">
            <input
              type="text"
              disabled
              class="number"
              value={data.bestehende_arbeitsverhaeltnisse_brutto}
            />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>

      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={54} />
        </div>
        <div class="content">
          <div>
            Darin ist ein Arbeitgeberanteil zu vermögenswirksamen Leistungen
            enthalten
          </div>
          <div class="user-value">
            <select
              bind:value={
                data.bestehende_arbeitsverhaeltnisse_arbeitgeberanteil_vermoegenswirksame_leistungen
              }
            >
              <option value={false}>Nein</option>
              <option value={true}>Ja</option>
            </select>
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>

      {#if data.bestehende_arbeitsverhaeltnisse_arbeitgeberanteil_vermoegenswirksame_leistungen}
        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={55} />
          </div>
          <div class="content">
            <div>
              Ausbildungs- und Praktikumsvergütung brutto – auch Sachbezüge
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.ausbildungs_und_praktikumsverguetung_brutto}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={56} />
          </div>
          <div class="content">
            <div>
              Einkünfte aus selbstständiger Arbeit, Gewerbebetrieb, Vermietung
              und Verpachtung, Land- und Forstwirtschaft
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.einkuenfte_aus_selbststaendiger_arbeit_gewerbe}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={57} />
          </div>
          <div class="content">
            <div>
              Einnahmen aus Kapitalvermögen (z. B. Sparzinsen) vor Abzug des
              Sparerpauschbetrages
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.einnahmen_aus_kapitalvermoegen}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={58} />
          </div>
          <div class="content">
            <div>
              Waisenrente und/oder Waisengeld (einschl. Weihnachtszuwendung),
              sonstige Renten (z. B. Unfallrenten)
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.waisenrente_und_sonstige_renten}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={59} />
          </div>
          <div class="content">
            <div>
              Ausbildungsbeihilfen und gleichartige Leistungen sowie Leistungen
              nach dem AFBG
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.ausbildungsbeihilfen_und_gleichartige_leistungen}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={60} />
          </div>
          <div class="content">
            <div>
              Unterhaltsleistungen (nicht der Eltern), die für mich bestimmt
              sind
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.unterhaltsleistungen_nicht_der_eltern}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={61} />
          </div>
          <div class="content">
            <div>
              Unterhaltsvorschussleistungen nach dem Unterhaltsvorschussgesetz
              (UhVorschG)
            </div>
            <div class="user-value">
              <input
                type="text"
                disabled
                class="number"
                value={data.unterhalts_vorschussleistungen}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>

        <div class="list-entry">
          <div class="number">
            <ApplicationIndicator indicator={62} />
          </div>
          <div class="content">
            <div>Weitere Einnahmen (nicht: laufende BAföGZahlungen)</div>
            <div class="user-value">
              <input
                type="text"
                class="number"
                disabled
                value={data.weitere_einnahmen}
              />
              <!-- <textarea disabled>{content}</textarea> -->
            </div>
            <div class="notes">
              <ApplicationComments />
            </div>
          </div>
        </div>
      {/if}

      <!-- {#each [1] as item (item)}
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={item + 30} />
        </div>
        <div class="content">
          <div>Aktuelle Gültigkeit von</div>
          <div>
            <textarea disabled>{content}</textarea>
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
    {/each} -->
    </div>
  </div>

  <div class="own-capital">
    <h3>Eigenkapital</h3>
    <div class="list">
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={63} />
        </div>
        <div class="content">
          <div>
            Bei Antragstellung bin ich <strong>unter 30 Jahre</strong> alt und meine
            Vermögenswerte der nachfolgenden Positionen in den Zeilen 65 bis 75 betragen
            insgesamt weniger als 10.000 Euro.
          </div>
          <div class="user-value">
            <label>
              <span>Ja</span>
              <input
                type="radio"
                class="number"
                name="under_30_and_under_10k_capital"
                disabled
              />
            </label>
            <label>
              <span>Nein</span>
              <input
                type="radio"
                class="number"
                name="under_30_and_under_10k_capital"
                disabled
              />
            </label>
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={64} />
        </div>
        <div class="content">
          <div>
            Bei Antragstellung bin ich <strong>über 30 Jahre</strong> alt und meine
            Vermögenswerte der nachfolgenden Positionen in den Zeilen 65 bis 75 betragen
            insgesamt weniger als 30.000 Euro.
          </div>
          <div class="user-value">
            <label>
              <span>Ja</span>
              <input
                type="radio"
                class="number"
                name="over_30_and_under_30k_capital"
                disabled
              />
            </label>
            <label>
              <span>Nein</span>
              <input
                type="radio"
                class="number"
                name="over_30_and_under_30k_capital"
                disabled
              />
            </label>
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={65} />
        </div>
        <div class="content">
          <div>Höhe des Barvermögens (Bargeld)</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={66} />
        </div>
        <div class="content">
          <div>
            Höhe der Bank- und Sparguthaben, einschließlich der Guthaben auf
            Girokonten und OnlineKonten
          </div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={67} />
        </div>
        <div class="content">
          <div>Höhe der Bauspar- und Prämienspargut haben</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={68} />
        </div>
        <div class="content">
          <div>
            Wertpapiere (z. B. Aktien, Pfandbriefe, Schatzanweisungen, Wechsel,
            Schecks)
          </div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={69} />
        </div>
        <div class="content">
          <div>Kraftfahrzeuge</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={70} />
        </div>
        <div class="content">
          <div>Lebensversicherungen</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={71} />
        </div>
        <div class="content">
          <div>
            Höhe von steuerlich gefördertem Alters vorsorge vermögen
            („Riester-Rente“)
          </div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={72} />
        </div>
        <div class="content">
          <div>
            Grundstücke, Häuser, Eigentumswohnungen (auch Miteigentumsanteile)
          </div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={73} />
        </div>
        <div class="content">
          <div>Betriebsvermögen (auch Miteigentumsanteile)</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={74} />
        </div>
        <div class="content">
          <div>
            Geldforderungen, digitales Vermögen (z. B. Kryptowährungen) und
            sonstige Rechte
          </div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
      <div class="list-entry">
        <div class="number">
          <ApplicationIndicator indicator={75} />
        </div>
        <div class="content">
          <div>Sonstige Vermögensgegenstände</div>
          <div class="user-value">
            <input type="text" class="number" disabled />
            <!-- <textarea disabled>{content}</textarea> -->
          </div>
          <div class="notes">
            <ApplicationComments />
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="sass">

  .calculations
    margin-bottom: 2rem
    
    h3
      margin-bottom: 1rem
    
    table
      width: 100%
      border-collapse: collapse
      background-color: var(--background-color-secondary)
      border-radius: 8px
      overflow: hidden
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1)
      
      tbody
        tr
          border-bottom: 1px solid var(--border-color, rgba(255, 255, 255, 0.1))
          
          &:last-child
            border-bottom: none
            
          &.sub-result
            background-color: var(--background-color-tertiary, rgba(255, 255, 255, 0.05))
            font-weight: 500
            
          &.result
            background-color: var(--accent-color, #007acc)
            color: white
            font-weight: 600
            
            td
              color: white
        
        td
          padding: 1rem
          color: var(--font-color)
          
          &:first-child
            font-weight: 500
            
          &:last-child
            text-align: right
            font-family: 'Courier New', monospace
            font-weight: 600

  .number
    display: flex

  .content
    padding-left: 1rem
    flex-grow: 1
    display: flex
    flex-direction: column
    gap: 0.5rem

  .list
    // display: flex
    // flex-direction: column
    display: grid
    grid-template-columns: 1fr 1fr
    gap: 2rem
    align-items: stretch

    .list-entry
      display: flex
      flex-direction: row
      gap: .5rem

      .content
        height: 100%
        display: flex
        flex-direction: column
        .user-value
          margin-top: auto
          padding-top: 1rem
          display: flex
          justify-content: flex-end
          textarea, input, select
            max-width: 100%
            width: 100%
            border-box: border-box
            field-sizing: content
            color: var(--font-color)
            resize: none
            border-radius: 5px
            border: none
            padding: 1rem
            background-color: var(--background-color-secondary)
            font-size: 1rem
            &.number
              text-align: right


        .notes
          summary
            div
              display: inline-flex
              align-items: center
              user-select: none
              cursor: pointer
              span:first-of-type
                flex-grow: 1


</style>
