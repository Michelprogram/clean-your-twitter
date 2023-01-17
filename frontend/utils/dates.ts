const optionsFull: Intl.DateTimeFormatOptions = {
  hour: "numeric",
  minute: "numeric",
  month: "long",
  year: "numeric",
  day: "numeric",
};

const optionsShort: Intl.DateTimeFormatOptions = {
  month: "long",
  year: "numeric",
  day: "numeric",
};

export default class Formatter {
  static full = Intl.DateTimeFormat("en-US", optionsFull);
  static short = Intl.DateTimeFormat("en-US", optionsShort);
}
