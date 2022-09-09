# Template: User Guide  

This template is for Crunchy Data employees to use as a reference when creating a User Guide. The easiest way to use this is to download the repository and import the files/content into your own project.

We use [Hugo](https://gohugo.io/getting-started/installing/) to generate docs pages.

> Use version [<= 0.102](https://github.com/gohugoio/hugo/releases). Crunchy docs currently build with Hugo 0.102.3.

## How to use

Assuming that your project is initialized as a git repository:

1. Download this repository in ZIP format and extract the contents into your project repository.
    - If your project repository contains other source code, the contents of priv-all-doc-userguide-template should go into a `docs` subdirectory.
2. Add the [Crunchy Hugo theme](https://github.com/CrunchyData/crunchy-hugo-theme) as a git submodule:

```
git submodule add https://github.com/CrunchyData/crunchy-hugo-theme themes/crunchy-hugo-theme
```
This will create a new `themes` subdirectory.  
3. Add content in the `content` [subdirectory](./content/). Images and other assets go under the `static` subdirectory.

To learn more, check out this page on the Hugo [directory structure](https://gohugo.io/getting-started/directory-structure/).

## How to test

In the user guide root directory, run the Hugo server:

```
hugo server
```
Then, go to localhost:1313 in your browser.

## Before You Write

Ask yourself:

- Who is going to read this guide? (Is there any chance that they will be a complete beginner?)
- What does the reader need to know for everyday usage? How do you make it easy for them to find?
