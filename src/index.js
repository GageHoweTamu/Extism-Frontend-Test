import createPlugin from '@extism/extism';

async function runPlugin() {
  const plugin = await createPlugin(
    'https://cdn.modsurfer.dylibso.com/api/v1/module/be716369b7332148771e3cd6376d688dfe7ee7dd503cbc43d2550d76cb45a01d.wasm',
    { useWasi: true }
  );
  
  const input = "Hello World";
  let out = await plugin.call("count_vowels", input);
  console.log(out.text());
  // => {"count": 3, "total": 3, "vowels": "aeiouAEIOU"}
}

runPlugin();