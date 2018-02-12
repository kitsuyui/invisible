# Invisible

Embed noise or hidden message to text.

## Add noise

```console
$ invisible add-noise < example/plain.txt > example/noised.txt
```

It looks normal. But the size is grown from 446 bytes to 1070 bytes by noise.

```
L​o⁡re⁣m ⁠ip⁠su⁡m‍ ⁠do​lor ‍sit ‌a‌met,⁠ ​c⁡on⁡sec⁠t⁣etur a⁡d‌i‌p‌is‌ci⁡ng⁣ ⁤elit⁤,⁢ ⁠se‌d⁤ ⁢d‌o ​e​i⁣us​mod ‍tem⁠por inc‌i⁡di​du⁡nt​ u⁢t la⁠bo⁤re et‌ ‍do‌lo⁢re‌ m⁤a⁣gna​ a⁡l‍i⁢qua.⁣ Ut ⁤enim ad‌ ‍min​im ‌ve⁠ni⁣am⁣,⁣ q⁠uis ⁢n‍o‌s⁢tru‌d ⁤e​x‍e⁢r‍c⁠i⁣ta‌ti⁡on⁡ ull⁣a⁤m‍c⁠o⁢ l‍abo⁡ris⁠ ⁢n‌isi​ ut⁢ a⁠l​i⁤q⁢u⁠ip e⁠x⁠ ⁠e⁡a⁢ co⁢mm⁣od⁤o ⁠conse‌q⁠u‌at. ⁣Du⁠i⁢s ‌aut⁤e ir⁠u​re do‍lor ​in⁠ r⁣ep⁢re⁤h⁡e⁢nderit i⁤n ‍v⁣o‌l⁣up​t​a⁠t⁠e v‌eli​t ⁡e‌sse c​illu‍m ‌do⁠l‌or‍e e​u⁡ ⁠fu​giat‍ n⁠ul⁣la​ ‍pa⁢r⁠i⁣at⁤u⁡r. ⁢Excep⁣teu‍r sin⁢t⁢ ‍o⁢c‍ca⁣e‌cat⁤ ⁣cupi⁣da⁤ta⁠t‍ ⁤n⁢on ⁤proident⁢, ⁠su⁢nt i⁠n⁢ ⁢culp​a‌ qui⁠ off‌ici‍a⁢ de⁠s⁠e⁠run⁤t m⁣ol⁠li​t ⁣an⁣im i⁠d⁣ e​st ‍l‌ab⁤orum⁡.
```


## Encode

```console
$ invisible encode -m 'Hello, World!' < example/plain.txt > example/embeded.txt
```

## Encoded text

It looks normal. But the hidden message is in.

```
L‍o‍r​e⁣m‍ ⁢i⁢p⁡s⁠u⁠m​ ⁣d⁤o⁡l⁢o⁡r‌ ​s​i⁢t⁠ ⁢a⁢m⁤e⁠t⁡,⁡ ⁣c⁣o‌n⁡s⁡e‌c​t‍e​t​u​r​ ​a‍d‍i​p⁣i‍s⁢c⁢i⁡n⁠g⁠ ​e⁣l⁤i⁡t⁢,⁡ ‌s​e​d⁢ ⁠d⁢o⁢ ⁤e⁠i⁡u⁡s⁣m⁣o‌d⁡ ⁡t‌e​m‍p​o​r​ ​i​n‍c‍i​d⁣i‍d⁢u⁢n⁡t⁠ ⁠u​t⁣ ⁤l⁡a⁢b⁡o‌r​e​ ⁢e⁠t⁢ ⁢d⁤o⁠l⁡o⁡r⁣e⁣ ‌m⁡a⁡g‌n​a‍ ​a​l​i​q​u‍a‍.​ ⁣U‍t⁢ ⁢e⁡n⁠i⁠m​ ⁣a⁤d⁡ ⁢m⁡i‌n​i​m⁢ ⁠v⁢e⁢n⁤i⁠a⁡m⁡,⁣ ⁣q‌u⁡i⁡s‌ ​n‍o​s​t​r​u​d‍ ‍e​x⁣e‍r⁢c⁢i⁡t⁠a⁠t​i⁣o⁤n⁡ ⁢u⁡l‌l​a​m⁢c⁠o⁢ ⁢l⁤a⁠b⁡o⁡r⁣i⁣s‌ ⁡n⁡i‌s​i‍ ​u​t​ ​a​l‍i‍q​u⁣i‍p⁢ ⁢e⁡x⁠ ⁠e​a⁣ ⁤c⁡o⁢m⁡m‌o​d​o⁢ ⁠c⁢o⁢n⁤s⁠e⁡q⁡u⁣a⁣t‌.⁡ ⁡D‌u​i‍s​ ​a​u​t​e‍ ‍i​r⁣u‍r⁢e⁢ ⁡d⁠o⁠l​o⁣r⁤ ⁡i⁢n⁡ ‌r​e​p⁢r⁠e⁢h⁢e⁤n⁠d⁡e⁡r⁣i⁣t‌ ⁡i⁡n‌ ​v‍o​l​u​p​t​a‍t‍e​ ⁣v‍e⁢l⁢i⁡t⁠ ⁠e​s⁣s⁤e⁡ ⁢c⁡i‌l​l​u⁢m⁠ ⁢d⁢o⁤l⁠o⁡r⁡e⁣ ⁣e‌u⁡ ⁡f‌u​g‍i​a​t​ ​n​u‍l‍l​a⁣ ‍p⁢a⁢r⁡i⁠a⁠t​u⁣r⁤.⁡ ⁢E⁡x‌c​e​p⁢t⁠e⁢u⁢r⁤ ⁠s⁡i⁡n⁣t⁣ ‌o⁡c⁡c‌a​e‍c​a​t​ ​c​u‍p‍i​d⁣a‍t⁢a⁢t⁡ ⁠n⁠o​n⁣ ⁤p⁡r⁢o⁡i‌d​e​n⁢t⁠,⁢ ⁢s⁤u⁠n⁡t⁡ ⁣i⁣n‌ ⁡c⁡u‌l​p‍a​ ​q​u​i​ ‍o‍f​f⁣i‍c⁢i⁢a⁡ ⁠d⁠e​s⁣e⁤r⁡u⁢n⁡t‌ ​m​o⁢l⁠l⁢i⁢t⁤ ⁠a⁡n⁡i⁣m⁣ ‌i⁡d⁡ ‌e​s‍t​ ​l​a​b​o‍r‍u​m⁣.‍
```

## Decode

```console
$ invisible decode < example/embeded.txt
Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!
```

## LICENSE

### Source

The 3-Clause BSD License. See also LISENCE file.
