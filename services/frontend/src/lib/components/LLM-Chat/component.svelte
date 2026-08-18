<script lang="ts">
  import { requestRagSchueler } from "$lib/api/rag/request-rag-schueler";
  import { requestRagStudierenden } from "$lib/api/rag/request-rag-studierenden";
  import { startConversationSchueler } from "$lib/api/rag/start-conversation-schueler";
  import type { Message } from ".";
  import { Button } from "../Button";
  import { compile } from "mdsvex";

  let ragInputValue: string = $state("");

  interface Props {
    ragConversationID?: string | undefined;
    ragDataSource: "schueler" | "studierenden";
  }

  let { ragConversationID = $bindable(), ragDataSource }: Props = $props();

  let llmMessages: Message[] = $state([]);
  let isGeneratingResponse: boolean = $state(false);
  let abortController: AbortController | undefined = $state(undefined);

  async function newLLMSession() {
    abortController?.abort();
    abortController = undefined;
    isGeneratingResponse = false;
    ragConversationID = undefined;
    llmMessages = [];
  }

  function stopLLMResponse() {
    abortController?.abort();
    abortController = undefined;
    isGeneratingResponse = false;
  }

  async function sendLLMChatMessage() {
    let messageContent = ragInputValue;
    ragInputValue = "";

    llmMessages.push({
      message: messageContent,
      html: await parseMarkdownToHTML(messageContent),
      direction: "input",
    });

    if (ragConversationID === undefined) {
      const conversation = await startConversationSchueler();
      ragConversationID = conversation?.ID;
    }

    abortController?.abort();
    abortController = new AbortController();
    isGeneratingResponse = true;

    try {
      if (ragDataSource === "schueler") {
        await requestRagSchueler(
          ragConversationID,
          messageContent,
          abortController.signal,
        ).then(handleRagResponseStreamasync);
      } else if (ragDataSource === "studierenden") {
        await requestRagStudierenden(
          ragConversationID,
          messageContent,
          abortController.signal,
        ).then(handleRagResponseStreamasync);
      }
    } finally {
      isGeneratingResponse = false;
      abortController = undefined;
    }
  }

  async function handleRagResponseStreamasync(
    reader:
      | {
          read: () =>
            | PromiseLike<{ value: any; done: any }>
            | { value: any; done: any };
        }
      | undefined,
  ) {
    if (reader !== undefined) {
      const index = llmMessages.push({
        message: "",
        html: "",
        direction: "output",
      });

      const decoder = new TextDecoder();
      let done = false;

      while (!done) {
        const { value, done: doneReading } = await reader.read();

        done = doneReading;
        if (value) {
          llmMessages[index - 1].message += decoder.decode(value);
          llmMessages[index - 1].html = await parseMarkdownToHTML(
            llmMessages[index - 1].message,
          );
        }
      }
    }
  }

  $effect(() => {
    return () => {
      abortController?.abort();
      abortController = undefined;
    };
  });

  async function parseMarkdownToHTML(input: string): Promise<string> {
    return (await compile(input, {}))?.code ?? "";
  }
</script>

<div class="llm-chat">
  <div class="info-header">
    <div class="danger danger-legaly">
      <strong>Hinweis:</strong> Diese Funktion befindet sich derzeit in der Beta-Phase.
      Bitte beachten Sie, dass die generierten Antworten ungenau sein können und
      nicht als offizielle Auskunft betrachtet werden sollten!
    </div>
  </div>
  {#if llmMessages.length === 0}
    <div class="no-conversation">
      <div class="text">
        <span class="header">Lass uns loslegen!</span>
        <span
          >Schreibe einen Propmt, um Inhalte aus den hochgeladenen Dokumenten zu
          überprüfen</span
        >
      </div>
    </div>
  {:else}
    <ul class="history">
      {#each llmMessages as llmMessage}
        <li class="message message-{llmMessage.direction}">
          <div class="message-body">
            <div class="message-content">
              <span>{@html llmMessage.html}</span>
            </div>
          </div>
        </li>
      {/each}
    </ul>
  {/if}
  {#if isGeneratingResponse}
    <div class="processing-indicator">
      <span>Server is processing...</span>
      <Button onclick={() => stopLLMResponse()}>Abbruch</Button>
    </div>
  {/if}
  <div class="inputs">
    <textarea bind:value={ragInputValue} placeholder="Schreibe deinen Prompt"
    ></textarea>
    <Button onclick={() => sendLLMChatMessage()}>Send</Button>
    <Button onclick={() => newLLMSession()}>Neue Session</Button>
  </div>
</div>

<style lang="sass">
  .llm-chat
    width: 100%
    height: 100%
    display: flex
    flex-direction: column
    background-color: var(--background-color)

    .info-header
      .danger
        border-left: 5px solid var(--color-red)
        background-color: var(--color-red-10)
        padding: 1rem

    .no-conversation
      width: 100%
      min-height: 20vh
      position: relative
      .text
        transform: translate(-50%,-50%)
        left: 50%
        top: 50%
        position: absolute
        display: flex
        flex-direction: column
        gap: 1rem
        text-align: center
        span.header
            font-weight: bolder
            font-size: 2rem


    .history
      display: block
      flex-grow: 1
      display: flex
      flex-direction: column
      gap: 2rem
      padding: 1rem 2rem
      max-height: 70vh
      overflow-y: scroll

      .message
        display: flex
        list-style-type: none

        &.message-input
          justify-content: end
          margin-left: 2rem

          .message-content
            background-color: var(--color-blue)
            color: var(--font-color-white)

        &.message-output
          margin-right: 2rem
          .message-content
            background-color: var(--background-color-tertiary)
            color: var(--font-color-white)

        .message-content
          display: inline-block
          padding: 1rem
          border-radius: 5px

    .processing-indicator
      margin-top: auto
      display: flex
      align-items: center
      gap: 1rem
      padding: 0.5rem 0

    .inputs
      margin-top: auto
      display: flex
      gap: 1rem

      textarea
        width: 100%
        padding: .5rem 1rem
        font-size: 1rem
        display: block
        box-sizing: border-box
</style>
