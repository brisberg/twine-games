window.pushEvent = (category, action, label, value) => {
  window.dataLayer.push({
    event: GAME_NAME,
    category,
    action,
    label,
    value,
  });
};
