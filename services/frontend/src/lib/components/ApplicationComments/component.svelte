<script lang="ts">
  import { Accordion } from "bits-ui";
</script>

<div class="dcs-theme">
  <Accordion.Root type="single">
    <Accordion.Item>
      <Accordion.Header>
        <Accordion.Trigger class="trigger">
          <!-- <span>Weitere Informationen</span> -->
          <span></span>
          <span>0 Kommentare</span>
        </Accordion.Trigger>
      </Accordion.Header>
      <Accordion.Content>
        <div class="comments">
          <h3>Kommentare</h3>
          <ul class="comments-list">
            <li>
              <div class="comment-entry">
                <div class="comment-author">
                  <span class="author-name">Max Mustermann</span>
                  <span class="comment-date">01.01.2024</span>
                </div>
                <div class="comment-content">
                  <p>
                    Lorem ipsum dolor, sit amet consectetur adipisicing elit.
                    Harum maxime dignissimos, dolore dolores veritatis atque
                    magnam dolorum cumque natus quos ducimus commodi.
                    Repellendus quas quam reiciendis, mollitia consequuntur sunt
                    aspernatur!
                  </p>
                </div>
                <div class="comment-actions">
                  <button class="reply-button">Antworten</button>
                  <button class="edit-button">Bearbeiten</button>
                  <button class="delete-button">Löschen</button>
                </div>
              </div>
            </li>
            <li>
              <div class="comment-entry">
                <div class="comment-author">
                  <span class="author-name">Max Mustermann</span>
                  <span class="comment-date">01.01.2024</span>
                </div>
                <div class="comment-content">
                  <p>
                    Lorem ipsum, dolor sit amet consectetur adipisicing elit.
                    Accusantium quia quibusdam doloremque cupiditate molestias
                    quas dolores ipsam reprehenderit ex porro! Maxime voluptatem
                    blanditiis reprehenderit labore modi vero officiis saepe
                    reiciendis.
                  </p>
                </div>
                <div class="comment-actions">
                  <button class="reply-button">Antworten</button>
                  <button class="edit-button">Bearbeiten</button>
                  <button class="delete-button">Löschen</button>
                </div>
              </div>
            </li>
            <li>
              <div class="comment-entry">
                <div class="comment-author">
                  <span class="author-name">Max Mustermann</span>
                  <span class="comment-date">01.01.2024</span>
                </div>
                <div class="comment-content">
                  <p>
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit.
                    Perferendis voluptate impedit necessitatibus, culpa itaque,
                    eveniet fuga numquam error quis magnam nobis debitis
                    excepturi nemo ex! Delectus aut temporibus modi sapiente.
                  </p>
                </div>
                <div class="comment-actions">
                  <button class="reply-button">Antworten</button>
                  <button class="edit-button">Bearbeiten</button>
                  <button class="delete-button">Löschen</button>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </Accordion.Content>
    </Accordion.Item>
  </Accordion.Root>
</div>

<style lang="sass">
  .dcs-theme
    .comments-list
      display: flex
      flex-direction: column
      gap: 2rem
      list-style: none
      padding: 0
    :global(.trigger)
      border: none
      background-color: transparent
      display: flex
      align-items: center
      justify-content: space-between
      width: 100%
      font-size: 1rem
      padding: 0
      user-select: none
      color: var(--font-color)

</style>
