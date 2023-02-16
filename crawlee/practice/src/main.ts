// For more information, see https://crawlee.dev/
import { PlaywrightCrawler, ProxyConfiguration, Dataset } from "crawlee";
// import { router } from "./routes.js";

const startUrls = ["https://crawlee.dev"];

const crawler = new PlaywrightCrawler({
  // proxyConfiguration: new ProxyConfiguration({ proxyUrls: ['...'] }),
  maxRequestsPerCrawl: 5,
  // - request: an instance of the Request class with information such as URL and HTTP method
  // - page: Playwright's Page object (see https://playwright.dev/docs/api/class-page)

  async requestHandler({ request, page, enqueueLinks, log }) {
    log.info(`Processing ${request.url}...`);

    const locator = page.getByRole("div");
    const data = await locator.evaluateAll((divs) => {
      const scrapedData: { value: HTMLDivElement | null }[] = [];

      divs.forEach((div) => {
        scrapedData.push({
          value: div.querySelector("div"),
        });
      });
      return scrapedData;
    });
    // A function to be evaluated by Playwright within the browser context.
    // const locate = page.locator("document");
    // const data = await locate.evaluateAll(($posts) => {
    //   const scrapedData: { value: HTMLBodyElement | null }[] = [];
    //   $posts.forEach(($post) => {
    //     scrapedData.push({
    //       value: $post.querySelector("body"),
    //       //   rank: $post.querySelector(".rank").innerText,
    //       //   href: $post.querySelector(".title a").href,
    //     });
    //   });

    //   return scrapedData;
    // });

    // Store the results to the default dataset.
    await Dataset.pushData({ data });

    // Find a link to the next page and enqueue it if it exists.
    const infos = await enqueueLinks({
      //   selector: ".morelink",
      strategy: "all",
    });

    if (infos.processedRequests.length === 0)
      log.info(`${request.url} is the last page!`);
  },

  // This function is called if the page processing failed more than maxRequestRetries+1 times.
  failedRequestHandler({ request, log }) {
    log.info(`Request ${request.url} failed too many times.`);
  },
});

// await crawler.addRequests([]);
await crawler.run(startUrls);

console.log("Crawler finished.");
