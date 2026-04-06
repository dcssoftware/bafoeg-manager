<script>
  import { ApplicationIndicator } from "$lib/components/ApplicationIndicator";
  import { IconArrowLeft, IconSquareCaretRight } from "$lib/components/Icons";
  import { page } from "$app/state";
  import { ApplicationComments } from "$lib/components/ApplicationComments";

  const applicationID = $state(page.params.applicationID);

  const content =
    "Lorem ipsum dolor sit amet consectetur adipisicing elit. Labore earum id iure voluptates velit voluptas itaque ex. Quidem, natus fugit. Commodi inventore quia adipisci aspernatur vero? Quisquam magni harum placeat.";
</script>

<a class="navigation-left" href="/applications/{applicationID}">
  <IconArrowLeft />
  <span>Zurück zum Antrag</span>
</a>

<h1>Diff</h1>

<div class="diff">
  <table>
    <thead>
      <tr>
        <th class="indicator"></th>
        <th>OLD</th>
        <th></th>
        <th>NEW</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="4">
          <span class="header">Allgemeine Informationen</span>
        </td>
      </tr>
      {#each new Array(5) as item, index}
        <tr>
          <td class="indicator" rowspan="2">
            <div class="indicator-element">
              <ApplicationIndicator indicator={index + 30} />
            </div>
          </td>
          <td class="description" colspan="3">
            <span> Aktuelle Gültigkeit von </span>
          </td>
        </tr>
        <tr class="content-diff">
          <td>
            <div class="value">
              <textarea disabled>{content}</textarea>
            </div>
          </td>
          <td class="arrow"><IconSquareCaretRight /></td>
          <td>
            <div class="value">
              <textarea disabled>{content}</textarea>
            </div>
          </td>
        </tr>
        <tr class="comments">
          <td></td>
          <td colspan="3">
            <ApplicationComments />
          </td>
        </tr>
      {/each}

      <tr>
        <td colspan="4">
          <span class="header">Eigene Finanzielle Situation</span>
        </td>
      </tr>
      {#each new Array(5) as item, index}
        <tr>
          <td class="indicator" rowspan="2">
            <div class="indicator-element">
              <ApplicationIndicator indicator={index + 30} />
            </div>
          </td>
          <td class="description" colspan="3">
            <span> Aktuelle Gültigkeit von </span>
          </td>
        </tr>
        <tr class="content-diff">
          <td>
            <div class="value">
              <textarea disabled>{content}</textarea>
            </div>
          </td>
          <td class="arrow"><IconSquareCaretRight /></td>
          <td>
            <div class="value">
              <textarea disabled>{content}</textarea>
            </div>
          </td>
        </tr>
        <tr class="comments">
          <td></td>
          <td colspan="3">
            <ApplicationComments />
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style lang="sass">
  .navigation-left
    :global(svg)
      $site: 1rem
      width: $site
      height: $site
      fill: var(--font-color)

  .diff
    background-color: var(--background-color-tertiary)
    padding: 2rem
    hr
      width: 90%
    table
      width: 100%
      border-collapse: collapse
      tr
        &.comments
          td
            padding: 1rem 1rem 2rem
        td.description
          padding: 1rem 1rem 1rem
        td.indicator
          width: 1%
          vertical-align: top
          .indicator-element
            display: flex
            align-items: center
            justify-content: center
        td.arrow
          $size: 1rem
          width: $size
          :global(svg)
            width: $size
            height: $size
            fill: var(--font-color)
        td
          .value
            display: flex
            flex-direction: row
            padding: 0 1rem
            textarea
              flex-grow: 1
    textarea
      border-box: border-box
      field-sizing: content
      color: var(--font-color)
      resize: none
      border-radius: 5px
      border: none
      padding: 1rem
      font-size: 1rem
      background-color: var(--background-color-secondary)
    .header
      $padding-width: .5rem 
      background-color: var(--background-color-secondary)
      padding: 0.5rem $padding-width
      font-weight: bold
      text-align: center
      margin: 2rem 0 1rem
      width: calc(100% - ($padding-width * 2))
      display: block
</style>
