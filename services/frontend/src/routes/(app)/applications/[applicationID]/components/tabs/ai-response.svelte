<script>
  import { AiAvatarPicture } from "$lib/assets";
  import { Button } from "$lib/components/Button";
</script>

<h2 class="header">AI Response</h2>

<div class="content">
  <div class="pre-check">
    <p>
      Lorem ipsum dolor, sit amet consectetur adipisicing elit. Laborum itaque
      repudiandae temporibus quasi minus voluptates saepe labore, voluptas,
      ducimus eos tempora nam expedita debitis. Delectus odit modi voluptates
      neque maxime.
    </p>
    <details open class="legal-text">
      <summary>Bundes Sozialgesetzbuch §12</summary>
      <div>
        <ul>
          <li>
            (1) Lorem ipsum dolor sit amet consectetur adipisicing elit.
            Consectetur ducimus nesciunt illo reprehenderit dolor modi
            voluptates, quas nostrum dolores nisi facere excepturi mollitia
            suscipit? Inventore doloribus libero quas porro? Doloremque.
          </li>
          <li>
            <mark
              >(2) Lorem ipsum dolor sit amet consectetur adipisicing elit.
              Consectetur ducimus nesciunt illo reprehenderit dolor modi
              voluptates, quas nostrum dolores nisi facere excepturi mollitia
              suscipit? Inventore doloribus libero quas porro? Doloremque.</mark
            >
          </li>
          <li>
            (3) Lorem ipsum dolor sit amet consectetur adipisicing elit.
            Consectetur ducimus nesciunt illo reprehenderit dolor modi
            voluptates, quas nostrum dolores nisi facere excepturi mollitia
            suscipit? Inventore doloribus libero quas porro? Doloremque.
          </li>
        </ul>
      </div>
    </details>
  </div>

  <div class="chat">
    <div class="writebox">
      <textarea spellcheck="false" placeholder="Ask some questions to the ai"
      ></textarea>
      <Button>Senden</Button>
    </div>
    <div class="dialog">
      <div class="message">
        <div class="avatar">
          <img src={AiAvatarPicture} alt="Avatar AI-Agent" />
        </div>
        <div class="text">
          <p>
            Lorem ipsum dolor, sit amet consectetur adipisicing elit. Laborum
            itaque repudiandae temporibus quasi minus voluptates saepe labore,
            voluptas, ducimus eos tempora nam expedita debitis. Delectus odit
            modi voluptates neque maxime.
          </p>
        </div>
      </div>
      <div class="message">
        <div class="avatar">
          <img src="/api/v1/self/picture" alt="Avatar Tim" />
        </div>
        <div class="text">
          <p>
            Lorem ipsum dolor, sit amet consectetur adipisicing elit. Laborum
            itaque repudiandae temporibus quasi minus voluptates saepe labore,
            voluptas, ducimus eos tempora nam expedita debitis. Delectus odit
            modi voluptates neque maxime.
          </p>
        </div>
      </div>
    </div>
  </div>
</div>

<style lang="sass">
  .content
    display: flex
    flex-direction: column
    gap: 2rem

    .pre-check
      line-height: 1.5

      details.legal-text
        background-color: var(--background-color-secondary)
        padding: 1rem
        border-radius: 5px

        ul
          display: flex
          flex-direction: column
          gap: 0.5rem
          li
            line-height: 1.5
        mark
          background-color: var(--color-yellow-30)
          color: var(--font-color)
          text-decoration: underline

    .chat
      display: flex
      flex-direction: column
      gap: 1rem
      .writebox
        display: flex
        gap: 1rem
        textarea
          padding: 1rem
          width: 100%
          height: 1rem
          resize: none
          font-size: 1rem
          border-radius: 5px
    .dialog
      display: flex
      flex-direction: column
      gap: 1rem
      .message
        display: flex
        gap: 1rem
        .avatar
          $size: 3rem
          width: $size
          height: $size
          aspect-ratio: 1
          img
            object-fit: cover
            width: 100%
            height: 100%
            border-radius: 50%
        .text
          background-color: var(--background-color-secondary)
          padding: 1rem
          border-radius: 5px
          p
            font-size: 1rem
            margin: 0
            line-height: 1.5
</style>
