import Button from "../../../components/button";
import styles from "../../shared/style/modal.module.css";
import { useDeleteDisciplinesMutation } from "../api/useDeleteDisciplinesMutation";

export function DeleteModal({ deleteDiscipline, isOpen, original }) {
  const deleteDisciplineMutation = useDeleteDisciplinesMutation();
  const id = original.id;

  if (!isOpen) {
    return null;
  }
  function deleteDisciplineHandle() {
    deleteDisciplineMutation.mutateAsync({ id });
    deleteDiscipline(true);
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <p className={styles['delete-text']}>Вы уверены, что хотите удалить {original.name}?</p>
        <div className={styles["modal-delete"]}>
          <Button
            styleFeature='close'
            type="submit"
            onClick={() => deleteDisciplineHandle()}
          >
            Удалить
          </Button>
          <Button onClick={deleteDiscipline}>Отмена</Button>
        </div>
      </div>
    </div>
  );
}
