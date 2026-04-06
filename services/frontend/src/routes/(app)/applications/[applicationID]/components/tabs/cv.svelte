<h2 class="header">Lebenslauf</h2>

<div class="cv-table">
  <table>
    <thead>
      <tr>
        <th>Von</th>
        <th>Bis</th>
        <th
          >Name und Ort der Ausbildungsstätte, des Arbeitgebers/Leistungsträgers</th
        >
        <th>Schulart/Fachrichtung; Art der Tätigkeit</th>
        <th
          >Erreichter Abschluss bzw. Bruttolohn / Höhe der Leistung monatlich in
          Euro</th
        >
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>08/12</td>
        <td>08/18</td>
        <td>Realschule, Musterstadt</td>
        <td>Realschule</td>
        <td>Realschulabschluss</td>
      </tr>
      <tr>
        <td>09/18</td>
        <td>08/21</td>
        <td>Fa. Mustermann & Co., Musterstadt</td>
        <td>Berufsausbildung zum Schlosser</td>
        <td>Gesellenbrief/Facharbeiterbrief, 15.08.2021</td>
      </tr>
      <tr>
        <td>09/21</td>
        <td>08/22</td>
        <td>Fa. Mustermann & Co., Musterstadt</td>
        <td>Schlosser</td>
        <td>2.300 Euro</td>
      </tr>
      <tr>
        <td>09/21</td>
        <td>08/22</td>
        <td>Fa. Mustermann & Co., Musterstadt</td>
        <td>Schlosser</td>
        <td>2.300 Euro</td>
      </tr>
      <tr>
        <td>09/22</td>
        <td>06/24</td>
        <td>Fachoberschule/Berufskolleg, Musterstadt</td>
        <td>Technischer Zweig</td>
        <td>Fachhochschulreife, 16.06.2024</td>
      </tr>
      <tr>
        <td>07/24</td>
        <td>03/25</td>
        <td>Fa. Müller & Co., Musterhausen</td>
        <td>Aushilfstätigkeit</td>
        <td>900 Euro</td>
      </tr>
      <tr>
        <td>04/25</td>
        <td>06/25</td>
        <td>Hochschule Musterhausen</td>
        <td>BWL, Bachelor</td>
        <td>ohne Abschluss</td>
      </tr>
      <tr>
        <td>07/25</td>
        <td>08/25</td>
        <td>Zwischenzeit</td>
        <td></td>
        <td></td>
      </tr>
      <tr>
        <td>09/25</td>
        <td></td>
        <td>Hochschule, Musterstadt</td>
        <td>Elektrotechnik, BA</td>
        <td></td>
      </tr>
    </tbody>
  </table>
</div>

<style lang="sass">
  .header
    margin-bottom: 1.5rem
    color: var(--color-text-primary)
    font-size: 1.5rem
    font-weight: 600

  .cv-table
    width: 100%
    overflow-x: auto
    border-radius: 0.5rem
    box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1)
    background-color: var(--background-color-primary)

    table
      width: 100%
      border-collapse: collapse
      font-size: 0.875rem
      line-height: 1.4

      th
        padding: 1rem 0.75rem
        text-align: left
        font-weight: 600
        font-size: 0.8125rem
        text-transform: uppercase
        letter-spacing: 0.025em
        color: var(--color-text-secondary)
        background-color: var(--color-gray-50)
        border-bottom: 2px solid var(--color-gray-200)
        position: sticky
        top: 0
        z-index: 10

      td
        padding: 0.875rem 0.75rem
        text-align: left
        color: var(--color-text-primary)
        background-color: var(--background-color-primary)
        border-bottom: 1px solid var(--color-gray-100)
        vertical-align: top

      tbody tr
        transition: background-color 0.15s ease-in-out

        &:hover
          background-color: var(--color-gray-25)

        &:last-child td
          border-bottom: none

      // Responsive column widths
      th:nth-child(1), td:nth-child(1)
        width: 8%
        min-width: 60px

      th:nth-child(2), td:nth-child(2)
        width: 8%
        min-width: 60px

      th:nth-child(3), td:nth-child(3)
        width: 30%
        min-width: 200px

      th:nth-child(4), td:nth-child(4)
        width: 25%
        min-width: 180px

      th:nth-child(5), td:nth-child(5)
        width: 29%
        min-width: 200px

    // Mobile responsiveness
    @media (max-width: 768px)
      font-size: 0.8125rem
      
      table th, table td
        padding: 0.625rem 0.5rem
        
      table th
        font-size: 0.75rem
</style>
