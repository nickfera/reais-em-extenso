const classesSingulares: string[] = [
  "mil",
  "milhão",
  "bilhão",
  "trilhão",
  "quatrilhão",
];

const classesPlurais: string[] = [
  "mil",
  "milhões",
  "bilhões",
  "trilhões",
  "quatrilhões",
];

const unidades = {
  1: "um",
  2: "dois",
  3: "três",
  4: "quatro",
  5: "cinco",
  6: "seis",
  7: "sete",
  8: "oito",
  9: "nove",
  10: "dez",
  11: "onze",
  12: "doze",
  13: "treze",
  14: "catorze",
  15: "quinze",
  16: "dezesseis",
  17: "dezessete",
  18: "dezoito",
  19: "dezenove",
};

const dezenas = {
  2: "vinte",
  3: "trinta",
  4: "quarenta",
  5: "cinquenta",
  6: "sessenta",
  7: "setenta",
  8: "oitenta",
  9: "noventa",
};

const centenas = {
  redondo1: "cem",
  1: "cento",
  2: "duzentos",
  3: "trezentos",
  4: "quatrocentos",
  5: "quinhentos",
  6: "seiscentos",
  7: "setecentos",
  8: "oitocentos",
  9: "novecentos",
};

/** Converte um número até 999 em extenso. */
function parteEmExtenso(num: number): string {
  const palavras: string[] = [];

  let restAbs = Math.abs(num);

  if (restAbs === 100) {
    return centenas.redondo1;
  }

  if (restAbs > 100) {
    const centena = Math.trunc(restAbs / 100);

    palavras.push(centenas[centena]);

    restAbs -= centena * 100;
  }

  if (restAbs >= 20) {
    const dezena = Math.trunc(restAbs / 10);

    palavras.push(dezenas[dezena]);

    restAbs -= dezena * 10;
  }

  if (restAbs > 0) {
    palavras.push(unidades[restAbs]);
  }

  return palavras.join(" e ");
}

function extenso(num: number): string {
  if (num === 0) {
    return "zero";
  }

  const partesInt: string[] = [];

  const abs = Math.abs(num);
  const int = Math.trunc(abs);

  if (int > 0) {
    const intStr = int.toString();

    const nPartes = Math.ceil(intStr.length / 3);

    for (let i = 0; i < nPartes; i++) {
      const fim = intStr.length - 3 * i;
      const inicio = fim - 3;

      const parteNum = Number(intStr.slice(inicio > 0 ? inicio : 0, fim));

      if (parteNum === 0) {
        continue;
      }

      let parteExtenso = parteEmExtenso(parteNum);

      if (i === 1 && parteExtenso === "um") {
        parteExtenso = classesSingulares[0];
      } else if (i > 1 && parteExtenso === "um") {
        parteExtenso += ` ${classesSingulares[i - 1]}`;
      } else if (i >= 1) {
        parteExtenso += ` ${classesPlurais[i - 1]}`;
      }

      partesInt.unshift(parteExtenso);
    }

    if (num < 0) {
      partesInt.unshift("menos");
    }

    partesInt.push(int > 1 ? "reais" : "real");
  }

  let parteDec: string = "";

  const absStr = abs.toString();

  if (absStr.indexOf(".") >= 0) {
    let decStr = absStr.substring(absStr.indexOf(".") + 1);

    if (decStr.length > 2) {
      decStr = `${decStr.substring(0, 2)}.${decStr.substring(2)}`;
    }

    const dec = Math.round(+decStr);

    parteDec = `${parteEmExtenso(dec)} centavo${dec > 1 ? "s" : ""}`;

    if (int === 0) {
      parteDec += " de real";
    } else {
      parteDec = `e ${parteDec}`;
    }
  }

  return [...partesInt, parteDec].join(" ");
}
